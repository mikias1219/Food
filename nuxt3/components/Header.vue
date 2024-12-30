<template>
  <nav class="fixed top-12 left-0 w-full bg-green-500 p-4 md:p-6 z-10">
    <div class="container mx-auto flex justify-between items-center">
      <div class="relative group">
        <!-- Button to toggle dropdown -->
        <button 
          class="text-white bg-green-600 px-4 py-2 rounded hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-300"
        >
          Categories
        </button>
        
        <!-- Dropdown menu -->
        <ul 
          v-if="categories.length"
          class="absolute mt-2 bg-white text-green-700 shadow-lg rounded-lg py-2 w-48"
        >
          <li
            v-for="category in categories"
            :key="category.id"
            class="px-4 py-2 hover:bg-green-100 cursor-pointer"
            @click="handleCategoryClick(category.name)"
          >
            {{ category.name }}
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useQuery } from '@vue/apollo-composable';
import gql from 'graphql-tag';
import { ref, computed, watch } from 'vue';
import { useCategoryStore } from '@/stores/categoryStore'; // Assuming Pinia is used

const router = useRouter();

// GraphQL query to fetch categories
const GET_CATEGORIES = gql`
  query GetCategories {
    categories {
      id
      name
    }
  }
`;

// Fetch categories from the GraphQL server
const { result, loading, error } = useQuery(GET_CATEGORIES);

// Pinia store
const categoryStore = useCategoryStore();

// Watch the query result and update the store
watch(
  () => result.value?.categories,
  (newCategories) => {
    if (newCategories) {
      categoryStore.setCategories(newCategories);
    }
  },
  { immediate: true }
);

// Reactive categories list for this component
const categories = computed(() => result.value?.categories || []);

// Handle category clicks
const handleCategoryClick = (categoryName) => {
  console.log('Selected Category:', categoryName);
  categoryStore.setSelectedCategory(categoryName); // Assuming your store has a `setSelectedCategory` method.
  router.push({ name: 'CategoryPage', params: { category: categoryName } });
};
</script>

<style scoped>
/* Additional styling for dropdown */
.relative:hover ul {
  display: block;
}

ul {
  display: none;
}
</style>
