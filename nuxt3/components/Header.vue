<template>
  <nav class="fixed top-12 left-0 w-full bg-brown-500 p-4 md:p-6 z-10">
    <div class="container mx-auto flex justify-between items-center">
      <div class="relative">
        <!-- Button to toggle dropdown -->
        <button
          @click="toggleDropdown"
          class="text-white bg-brown-600 px-4 py-2 rounded flex items-center gap-2 hover:bg-brown-700 focus:outline-none focus:ring-2 focus:ring-brown-300"
        >
          Categories
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            :class="{'rotate-180': dropdownOpen, 'rotate-0': !dropdownOpen}"
            class="w-4 h-4 transition-transform duration-200"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </button>

        <!-- Dropdown menu -->
        <ul
          v-show="dropdownOpen"
          class="absolute mt-2 bg-white text-brown-700 shadow-lg rounded-lg py-2 w-40 text-sm transition-all duration-200 opacity-0"
          :class="{ 'opacity-100': dropdownOpen, 'opacity-0': !dropdownOpen }"
        >
          <li
            v-for="category in categories"
            :key="category.id"
            class="px-4 py-1 hover:bg-brown-100 cursor-pointer"
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
const dropdownOpen = ref(false);

// Function to toggle the dropdown
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

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
/* Styling for dropdown animations and muted brown color */
.nav {
  background-color: #8B4513; /* Muted brown */
}

.w-40 {
  width: 10rem; /* Smaller dropdown width */
}

.text-brown-700 {
  color: #5C4033;
}

.bg-brown-500 {
  background-color: #8B4513;
}

.bg-brown-600 {
  background-color: #704214;
}

.bg-brown-700 {
  background-color: #5C4033;
}

.focus:ring-brown-300 {
  box-shadow: 0 0 0 2px rgba(188, 143, 143, 0.5);
}

ul {
  transform-origin: top;
}

.rotate-180 {
  transform: rotate(180deg);
}

.rotate-0 {
  transform: rotate(0deg);
}
</style>
