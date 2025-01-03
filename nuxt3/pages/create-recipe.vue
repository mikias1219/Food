<template>
  <div class="container mx-auto py-10 px-6 max-w-4xl">
    <h1 class="text-4xl font-extrabold text-center mb-10 text-gray-800">Create a New Recipe</h1>

    <form @submit.prevent="submitRecipe" class="space-y-8 bg-white p-10 rounded-xl shadow-xl">
      <!-- Recipe Title -->
      <div>
        <label for="title" class="block text-lg font-semibold text-gray-700 mb-2">Title</label>
        <input
          v-model="form.title"
          id="title"
          type="text"
          class="w-full px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
          placeholder="Enter recipe title"
          required
        />
      </div>

      <!-- Recipe Description -->
      <div>
        <label for="description" class="block text-lg font-semibold text-gray-700 mb-2">Description</label>
        <textarea
          v-model="form.description"
          id="description"
          rows="4"
          class="w-full px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
          placeholder="Briefly describe the recipe"
          required
        ></textarea>
      </div>

      <!-- Category -->
      <div>
        <label for="category" class="block text-lg font-semibold text-gray-700 mb-2">Category</label>
        <select
          v-model="selectedCategory"
          id="category"
          class="w-full px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
          @change="handleCategoryChange"
          required
        >
          <option value="" disabled>Select a category</option>
          <option v-for="category in categories" :key="category.id" :value="category.id">
            {{ category.name }}
          </option>
          <option value="new">+ Add New Category</option>
        </select>
        <div v-if="isAddingCategory" class="mt-4 flex items-center gap-4">
          <input
            v-model="newCategory"
            type="text"
            class="flex-1 px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-600"
            placeholder="Enter new category name"
            required
          />
          <button
            @click.prevent="addNewCategory"
            class="bg-green-500 text-white px-6 py-2 rounded-md hover:bg-green-600 transition"
          >
            Add
          </button>
        </div>
      </div>

      <!-- Preparation Time -->
      <div>
        <label for="time" class="block text-lg font-semibold text-gray-700 mb-2">Preparation Time (minutes)</label>
        <input
          v-model="form.preparation_time"
          id="time"
          type="number"
          class="w-full px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
          placeholder="Enter preparation time"
          required
        />
      </div>

      <!-- Ingredients -->
      <div>
        <label class="block text-lg font-semibold text-gray-700 mb-2">Ingredients</label>
        <div v-for="(ingredient, index) in form.ingredients" :key="index" class="flex gap-4 items-center mb-4">
          <input
            v-model="ingredient.name"
            type="text"
            class="flex-1 px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            placeholder="Ingredient name"
            required
          />
          <input
            v-model="ingredient.quantity"
            type="text"
            class="flex-1 px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            placeholder="Quantity (e.g., 1 tsp)"
            required
          />
          <button
            @click.prevent="removeIngredient(index)"
            class="text-red-500 hover:text-red-600 transition"
          >
            Remove
          </button>
        </div>
        <button
          @click.prevent="addIngredient"
          class="text-blue-600 hover:text-blue-800 font-semibold transition"
        >
          + Add Ingredient
        </button>
      </div>

      <!-- Steps -->
      <div>
        <label class="block text-lg font-semibold text-gray-700 mb-2">Steps</label>
        <div v-for="(step, index) in form.steps" :key="index" class="flex gap-4 items-start mb-4">
          <input
            v-model="step.step_number"
            type="number"
            class="w-20 px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            placeholder="Step"
            required
          />
          <textarea
            v-model="step.description"
            rows="2"
            class="flex-1 px-4 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            placeholder="Step description"
            required
          ></textarea>
          <button
            @click.prevent="removeStep(index)"
            class="text-red-500 hover:text-red-600 transition"
          >
            Remove
          </button>
        </div>
        <button
          @click.prevent="addStep"
          class="text-blue-600 hover:text-blue-800 font-semibold transition"
        >
          + Add Step
        </button>
      </div>

      <!-- Recipe Image Upload -->
      <div>
    <label class="block text-lg font-semibold text-gray-700 mb-2">Upload Recipe Images</label>
    <div class="flex flex-col sm:flex-row gap-4 items-center">
      <label
        for="fileInput"
        class="bg-blue-500 text-white px-6 py-2 rounded-md cursor-pointer hover:bg-blue-600 transition"
      >
        Choose File
        <input
          id="fileInput"
          type="file"
          class="hidden"
          @change="handleFileChange"
        />
      </label>
      <p class="text-sm text-gray-500">Selected: {{ selectedFile?.name || 'None' }}</p>
      <button
        @click.prevent="uploadFile"
        :disabled="!selectedFile"
        class="bg-green-500 text-white px-6 py-2 rounded-md hover:bg-green-600 transition disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        Upload Image
      </button>
    </div>
    <p v-if="uploadStatus" class="text-sm text-gray-600 mt-2">{{ uploadStatus }}</p>
  </div>
      <!-- Submit Button -->
      <button
        type="submit"
        class="w-full bg-blue-500 text-white text-lg font-semibold px-6 py-3 rounded-md hover:bg-blue-600 transition"
      >
        Submit Recipe
      </button>

      <!-- Success and Error Messages -->
      <p v-if="successMessage" class="text-green-600 text-center mt-4">{{ successMessage }}</p>
      <p v-if="errorMessage" class="text-red-600 text-center mt-4">{{ errorMessage }}</p>
    </form>
  </div>
</template>



<script setup>
import { ref, onMounted } from "vue";
import { useApolloClient } from "@vue/apollo-composable";
import gql from "graphql-tag";

let recipeId = null;

const form = ref({
  title: "",
  description: "",
  category_id: "",
  user_id: "",
  preparation_time: "",
  ingredients: [{ name: "", quantity: "" }],
  steps: [{ step_number: 1, description: "" }],
  images: [{ image_url: "", is_featured: false }],
  shares: [{ is_shared: false }], // Adjusted for compatibility
});

const categories = ref([]);
const selectedCategory = ref("");
const newCategory = ref("");
const isAddingCategory = ref(false);

const successMessage = ref("");
const errorMessage = ref("");
const uploadStatus = ref("");
const selectedFile = ref(null);

const client = useApolloClient().client;

const fetchUserIdFromToken = () => {
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("token");
    if (token) {
      try {
        const decodedToken = JSON.parse(atob(token.split(".")[1]));
        form.value.user_id = decodedToken.userId;
      } catch (error) {
        console.error("Failed to decode token", error);
        errorMessage.value = "Invalid token. Please log in again.";
      }
    }
  }
};

const fetchCategories = async () => {
  try {
    const query = gql`
      query GetCategories {
        categories {
          id
          name
        }
      }
    `;
    const { data } = await client.query({ query, fetchPolicy: "no-cache" });
    categories.value = data.categories;
  } catch (error) {
    console.error("Failed to fetch categories:", error);
    errorMessage.value = "Failed to load categories. Please try again.";
  }
};

const handleCategoryChange = () => {
  isAddingCategory.value = selectedCategory.value === "new";
  form.value.category_id = isAddingCategory.value ? "" : selectedCategory.value;
};

const addNewCategory = async () => {
  try {
    const mutation = gql`
      mutation AddCategory($name: String!) {
        insert_categories_one(object: { name: $name }) {
          id
          name
        }
      }
    `;
    const { data } = await client.mutate({
      mutation,
      variables: { name: newCategory.value },
    });

    categories.value.push(data.insert_categories_one);
    selectedCategory.value = data.insert_categories_one.id;
    form.value.category_id = data.insert_categories_one.id;

    isAddingCategory.value = false;
    newCategory.value = "";
    successMessage.value = "Category added successfully!";
  } catch (error) {
    console.error("Failed to add new category:", error);
    errorMessage.value = "Failed to add new category. Please try again.";
  }
};

const addIngredient = () => {
  form.value.ingredients.push({ name: "", quantity: "" });
};

const removeIngredient = (index) => {
  form.value.ingredients.splice(index, 1);
};

const addStep = () => {
  const nextStepNumber = form.value.steps.length + 1;
  form.value.steps.push({ step_number: nextStepNumber, description: "" });
};

const removeStep = (index) => {
  form.value.steps.splice(index, 1);
  form.value.steps.forEach((step, i) => (step.step_number = i + 1));
};



const submitRecipe = async () => {
  if (!form.value.category_id) {
    errorMessage.value = "Please select or add a category.";
    return;
  }

  try {
    const mutation = gql`
      mutation AddRecipe($input: recipes_insert_input!) {
        insert_recipes_one(object: $input) {
          id
          title
        }
      }
    `;

    const variables = {
      input: {
        title: form.value.title,
        description: form.value.description,
        category_id: form.value.category_id,
        user_id: form.value.user_id,
        preparation_time: parseInt(form.value.preparation_time, 10),
        ingredients: {
          data: form.value.ingredients.map(({ name, quantity }) => ({ name, quantity })),
        },
        steps: {
          data: form.value.steps.map(({ step_number, description }) => ({
            step_number,
            description,
          })),
        },
        shares: {
          data: form.value.shares.map(({ is_shared }) => ({ is_shared })),
        },
      },
    };

    const { data } = await client.mutate({
      mutation,
      variables,
    });
    recipeId = data.insert_recipes_one.id
     successMessage.value = `Recipe "${data.insert_recipes_one.title}" added successfully!`;
    form.value = {
      title: "",
      description: "",
      category_id: "",
      user_id: "",
      preparation_time: "",
      ingredients: [{ name: "", quantity: "" }],
      steps: [{ step_number: 1, description: "" }],
      images: [{ image_url: "", is_featured: false }],
    
    };
  } catch (error) {
    console.error("Failed to submit recipe:", error);
    errorMessage.value = "Failed to submit recipe. Please try again.";
  }
};

// GraphQL mutation for file upload
const FILE_UPLOAD = gql`
  mutation fileUpload($name: String!, $type: String!, $base64str: String!,$RecipeID:String!) {
    fileUpload(name: $name, type: $type, base64str: $base64str,RecipeID:$RecipeID) {
      file_path
      RecipeID
    }
  }
`

const { mutate } = useMutation(FILE_UPLOAD)

const handleFileChange = async (event) => {
  const file = event.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      selectedFile.value = {
        name: file.name,
        type: file.type,
        base64str: e.target.result.split(',')[1],
      }
    }
    reader.readAsDataURL(file)
  }
}

const uploadFile = async () => {
  if (!selectedFile.value) {
    uploadStatus.value = 'No file selected!'
    return
  }

  try {
    uploadStatus.value = 'Uploading...'
    const { data } = await mutate({
      name: selectedFile.value.name,
      type: selectedFile.value.type,
      base64str: selectedFile.value.base64str,
      RecipeID:recipeId,
    })
    uploadStatus.value = `File uploaded successfully: ${data.fileUpload.file_path}`
  } catch (error) {
    uploadStatus.value = `Error: ${error.message}`
  }
}




onMounted(() => {
  fetchUserIdFromToken();
  fetchCategories();
});
</script>

