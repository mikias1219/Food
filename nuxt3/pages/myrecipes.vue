<template>
  <div class="container mx-auto py-8 px-4">
    <h1 class="text-3xl font-bold mb-6 text-center">My Recipes</h1>

    <!-- Loading and Error Messages -->
    <div v-if="loading" class="text-gray-500 text-center">Loading...</div>
    <div v-else-if="error" class="text-red-500 text-center">Error: {{ error.message }}</div>

    <!-- Recipe Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
      <div
        v-for="recipe in filteredRecipes"
        :key="recipe.id"
        class="border rounded-lg p-4 shadow-sm relative bg-white"
      >
        <img
          v-if="recipe.featured_image"
          :src="getImageUrl(recipe.featured_image)"
          alt="Recipe Image"
          class="w-full h-48 object-cover rounded-lg mb-4"
        />
        <h2 class="text-lg font-semibold mb-2">{{ recipe.title }}</h2>
        <p class="text-gray-600 mb-4">{{ recipe.description }}</p>
        <p class="text-sm text-gray-500">
          Prep Time: {{ recipe.preparation_time }} mins
        </p>
        <p class="text-sm text-gray-500">Category: {{ recipe.category.name }}</p>
        <div class="mt-4">
          <h3 class="font-semibold mb-1">Ingredients:</h3>
          <ul class="list-disc list-inside text-sm text-gray-700">
            <li v-for="ingredient in recipe.ingredients" :key="ingredient.id">
              {{ ingredient.quantity }} {{ ingredient.name }}
            </li>
          </ul>
        </div>
        <div class="mt-4">
          <h3 class="font-semibold mb-1">Steps:</h3>
          <ol class="list-decimal list-inside text-sm text-gray-700">
            <li v-for="step in recipe.steps" :key="step.id">
              {{ step.step_number }}. {{ step.description }}
            </li>
          </ol>
        </div>

        <!-- Buttons -->
        <div class="absolute top-2 right-2 flex flex-col gap-2">
          <button
            @click="startEditing(recipe)"
            class="px-3 py-1 bg-blue-500 text-white rounded-md text-sm"
          >
            Edit
          </button>
          <button
            @click="deleteRecipe(recipe.id)"
            class="px-3 py-1 bg-red-500 text-white rounded-md text-sm"
          >
            Delete
          </button>
          <button
            @click="toggleShareStatus(recipe.id)"
            class="px-3 py-1 bg-green-500 text-white rounded-md text-sm"
          >
            Share
          </button>
          <button
            @click="toggleUnshareStatus(recipe.id)"
            class="px-3 py-1 bg-gray-500 text-white rounded-md text-sm"
          >
            Unshare
          </button>
        </div>
      </div>
    </div>

    <!-- Edit Recipe Modal -->
    <div
      v-if="editingRecipe"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
    >
      <div class="bg-white p-6 rounded-lg w-full max-w-lg">
        <h2 class="text-xl font-bold mb-4">Edit Recipe</h2>
        <form @submit.prevent="updateRecipe">
          <div class="mb-4">
            <label class="block font-semibold mb-2">Title</label>
            <input
              v-model="editingRecipe.title"
              class="border rounded-md p-2 w-full"
              type="text"
            />
          </div>
          <div class="mb-4">
            <label class="block font-semibold mb-2">Description</label>
            <textarea
              v-model="editingRecipe.description"
              class="border rounded-md p-2 w-full"
            ></textarea>
          </div>
          <div class="mb-4">
            <label class="block font-semibold mb-2">Preparation Time</label>
            <input
              v-model="editingRecipe.preparation_time"
              class="border rounded-md p-2 w-full"
              type="number"
            />
          </div>
          <div class="mb-4">
            <h3 class="font-semibold">Ingredients</h3>
            <div
              v-for="(ingredient, index) in editingRecipe.ingredients"
              :key="ingredient.id || index"
              class="flex items-center gap-2"
            >
              <input
                v-model="ingredient.quantity"
                class="border rounded-md p-2 w-1/4"
                placeholder="Quantity"
                type="text"
              />
              <input
                v-model="ingredient.name"
                class="border rounded-md p-2 w-3/4"
                placeholder="Name"
                type="text"
              />
              <button
                type="button"
                @click="removeIngredient(index)"
                class="px-2 py-1 bg-red-500 text-white rounded-md"
              >
                Remove
              </button>
            </div>
            <button
              type="button"
              @click="addIngredient"
              class="mt-2 px-4 py-2 bg-green-500 text-white rounded-md"
            >
              Add Ingredient
            </button>
          </div>

          <div class="mb-4">
            <h3 class="font-semibold">Steps</h3>
            <div
              v-for="(step, index) in editingRecipe.steps"
              :key="step.id || index"
              class="flex items-center gap-2"
            >
              <input
                v-model="step.step_number"
                class="border rounded-md p-2 w-1/4"
                placeholder="Step Number"
                type="number"
              />
              <textarea
                v-model="step.description"
                class="border rounded-md p-2 w-3/4"
                placeholder="Description"
              ></textarea>
              <button
                type="button"
                @click="removeStep(index)"
                class="px-2 py-1 bg-red-500 text-white rounded-md"
              >
                Remove
              </button>
            </div>
            <button
              type="button"
              @click="addStep"
              class="mt-2 px-4 py-2 bg-green-500 text-white rounded-md"
            >
              Add Step
            </button>
          </div>

          <div class="flex gap-2">
            <button
              type="submit"
              class="px-4 py-2 bg-green-500 text-white rounded-md"
            >
              Save
            </button>
            <button
              @click="cancelEditing"
              type="button"
              class="px-4 py-2 bg-gray-500 text-white rounded-md"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>


<script setup>
import { ref,computed, onMounted } from 'vue';
import { gql } from '@apollo/client/core';
import { useNuxtApp } from '#app';
import { useRecipeStore } from '@/stores/recipe';  // Import the Pinia store
const recipeStore = useRecipeStore(); // Access the store

const userId = ref('');
const recipes = ref([]);
const editingRecipe = ref(null);
const loading = ref(true);
const error = ref(null);

const { $apolloClient } = useNuxtApp();
const searchQuery = ref('');
const filteredRecipes = computed(() => {
  return recipes.value.filter((recipe) =>
    recipe.title.toLowerCase().includes(recipeStore.searchQuery.toLowerCase())
  );
});

const backendBaseUrl = 'http://localhost:8085'; // Adjust according to your backend URL

const FETCH_RECIPES_QUERY = gql`
  query FetchRecipes($userId: uuid!) {
    recipes(where: { user_id: { _eq: $userId } }) {
      id
      title
      description
      preparation_time
      featured_image
      category {
        name
      }
      ingredients {
        id
        name
        quantity
      }
      steps {
        id
        step_number
        description
      }
    }
  }
`;

const DELETE_RECIPE_MUTATION = gql`
  mutation DeleteRecipe($id: uuid!) {
    delete_recipes_by_pk(id: $id) {
      id
    }
  }
`;

const UPDATE_RECIPE_MUTATION = gql`
  mutation UpdateRecipe($id: uuid!, $input: recipes_set_input!) {
    update_recipes_by_pk(pk_columns: { id: $id }, _set: $input) {
      id
      title
      description
      preparation_time
            ingredients {
        id
        name
        quantity
      }
      steps {
        id
        step_number
        description
      }
    }
  }
`;
const share_mutaion = gql`
  mutation shares($recipe_id: uuid!) {
  update_shares(
    where: { recipe_id: { _eq: $recipe_id } }
    _set: { is_shared: true }
  ) {
    affected_rows
  }
}
`;

const unshare_mutaion = gql`
  mutation unshare($recipe_id: uuid!) {
    update_shares(
      where: { recipe_id: { _eq: $recipe_id } }
      _set: { is_shared: false }
    ) {
      affected_rows
    }
  }
`;

const getImageUrl = (path) => {
  return path ? `${backendBaseUrl}${path}` : null;
};

const fetchRecipes = async () => {
  try {
    loading.value = true;
    const response = await $apolloClient.query({
      query: FETCH_RECIPES_QUERY,
      variables: { userId: userId.value },
    });
    recipes.value = response.data.recipes;
    error.value = null; // Clear any previous error
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
};

const deleteRecipe = async (id) => {
  if (confirm('Are you sure you want to delete this recipe?')) {
    try {
      await $apolloClient.mutate({
        mutation: DELETE_RECIPE_MUTATION,
        variables: { id },
      });
      recipes.value = recipes.value.filter((recipe) => recipe.id !== id);
    } catch (err) {
      alert('Failed to delete the recipe.');
    }
  }
};

const startEditing = (recipe) => {
  editingRecipe.value = { ...recipe };
};

const cancelEditing = () => {
  editingRecipe.value = null;
};

const updateRecipe = async () => {
  try {
    const { id, title, description, preparation_time } = editingRecipe.value;
    await $apolloClient.mutate({
      mutation: UPDATE_RECIPE_MUTATION,
      variables: { id, input: { title, description, preparation_time } },
    });
    fetchRecipes();
    editingRecipe.value = null;
  } catch (err) {
    alert('Failed to update the recipe.');
  }
};


const toggleShareStatus = async (shareid) => {
  console.log("Sharing recipe ID:", shareid);
  try {
    await $apolloClient.mutate({
      mutation: share_mutaion,
      variables: { recipe_id: shareid },
    });
    alert("Recipe shared successfully!");
  } catch (err) {
    alert("Failed to share the recipe.");
    console.error(err);
  }
};

const toggleUnshareStatus = async (recipeId) => {
  console.log("Unsharing recipe ID:", recipeId);
  try {
    await $apolloClient.mutate({
      mutation: unshare_mutaion,
      variables: { recipe_id: recipeId },
    });
    alert("Recipe unshared successfully!");
    // Optional: Refresh recipes to reflect changes
    await fetchRecipes();
  } catch (err) {
    alert("Failed to unshare the recipe.");
    console.error(err);
  }
};




onMounted(async () => {
  try {
    const token = localStorage.getItem('token');

    if (token) {
      const decodedToken = JSON.parse(atob(token.split('.')[1]));
      userId.value = decodedToken.userId;
    } else {
      throw new Error('Token not found. Please log in again.');
    }

    if (!userId.value) {
      throw new Error('User ID not found. Please log in again.');
    }

    await fetchRecipes();
  } catch (err) {
    error.value = err;
    console.error(err.message);
  } finally {
    loading.value = false;
  }
});
</script>

