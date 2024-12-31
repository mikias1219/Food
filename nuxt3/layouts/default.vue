<template>
  <div>
    <!-- Header with Dropdown Menus -->
    <header class="fixed top-0 left-0 w-full bg-green-500 shadow-md z-50">
      <nav class="container mx-auto px-6 py-4 flex items-center justify-between">
        <!-- Logo -->
        <a
          href="#"
          class="text-xl font-bold italic text-red-800 hover:text-blue-600"
          @click.prevent="navigateToHome"
        >
          FoodRecipes
        </a>

        <!-- Header Content -->
        <div v-if="!isLoginOrSignupPage" class="flex items-center space-x-6">
          <!-- Filter Section -->
          <div class="relative">
            <button
              @click="toggleFilter"
              class="flex items-center px-4 py-2 border border-gray-300 rounded-lg focus:outline-none hover:bg-gray-200"
            >
              Filter
              <span
                :class="{
                  'transform rotate-180': showFilter,
                  'transform rotate-0': !showFilter,
                }"
                class="ml-2 transition-transform duration-300"
              >
                ▼
              </span>
            </button>
            <transition name="fade">
              <div
                v-if="showFilter"
                class="absolute left-0 mt-2 w-48 bg-white border border-gray-300 rounded-lg shadow-lg z-10"
              >
                <ul class="py-2 text-sm text-gray-700">
                  <li v-for="filter in filters" :key="filter">
                    <button
                      class="w-full px-4 py-2 hover:bg-gray-100 text-left"
                      @click="applyFilter(filter)"
                    >
                      {{ filter }}
                    </button>
                  </li>
                </ul>
              </div>
            </transition>
          </div>

          <!-- Search Section -->
          <form class="relative flex-1" @submit.prevent="searchRecipes">
            <input
              v-model="recipeStore.searchQuery"
              type="text"
              placeholder="Search recipes..."
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
            />
            <button
              type="submit"
              class="absolute right-2 top-1/2 transform -translate-y-1/2 text-blue-500 hover:text-gray-800"
            >
              🔍
            </button>
          </form>

          <!-- Account Dropdown -->
          <div class="relative" v-if="authStore.isAuthenticated">
            <button
              class="flex items-center px-4 py-2 bg-gray-200 rounded-lg focus:outline-none hover:bg-gray-300"
              @click="toggleDropdown"
            >
              <img
                :src="authStore.user?.profilePicture || '/default-avatar.png'"
                alt="Profile"
                class="w-8 h-8 rounded-full mr-2"
              />
              {{ authStore.user?.name || 'My Account' }}
              <span
                :class="{
                  'transform rotate-180': showAccountDropdown,
                  'transform rotate-0': !showAccountDropdown,
                }"
                class="ml-2 transition-transform duration-300"
              >
                ▼
              </span>
            </button>
            <transition name="fade">
              <div
                v-if="showAccountDropdown"
                class="absolute right-0 mt-2 w-48 bg-white border border-gray-300 rounded-lg shadow-lg z-10"
              >
                <ul class="py-2 text-sm text-gray-700">
                  <li v-for="(action, key) in userActions" :key="key">
                    <button
                      class="w-full px-4 py-2 hover:bg-gray-100 text-left"
                      @click="navigateTo(action.route)"
                    >
                      {{ action.label }}
                    </button>
                  </li>
                  <li>
                    <button
                      class="w-full px-4 py-2 text-red-500 hover:bg-gray-100 text-left"
                      @click="logout"
                    >
                      Logout
                    </button>
                  </li>
                </ul>
              </div>
            </transition>
          </div>

          <!-- Login/Signup Buttons -->
          <div v-else>
            <button
              class="px-4 py-2 text-sm font-medium text-gray-600 hover:text-gray-800"
              @click="goToSignup"
            >
              Signup
            </button>
            <button
              class="px-4 py-2 text-sm font-medium text-white bg-blue-500 rounded-lg shadow hover:bg-blue-600"
              @click="goToLogin"
            >
              Login
            </button>
          </div>
        </div>
      </nav>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto py-24">
      <NuxtPage />
    </main>

    <footer class="bg-gray-800 text-white py-6">
      <div class="container mx-auto flex flex-col md:flex-row justify-between items-center">
        <p class="text-sm">&copy; 2024 FoodRecipes. All rights reserved.</p>
        <div class="flex space-x-4">
          <a href="#about" class="hover:text-blue-400">About Us</a>
          <a href="#top" class="hover:text-blue-400">Back to Top</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRecipeStore } from '@/stores/recipe';
import { useRouter, useRoute } from 'vue-router';

const authStore = useAuthStore();
const recipeStore = useRecipeStore();
const router = useRouter();
const route = useRoute();

const showFilter = ref(false);
const showAccountDropdown = ref(false);

const filters = ['Vegetarian', 'Non-Vegetarian', 'Vegan'];
const userActions = [
  { label: 'My Recipes', route: '/myrecipes' },
  { label: 'Create Recipe', route: '/create-recipe' },
  { label: 'My Profile', route: '/profile' },
];

const navigateToHome = () => router.push('/');
const goToSignup = () => router.push('/signup');
const goToLogin = () => router.push('/login');
const toggleFilter = () => (showFilter.value = !showFilter.value);
const toggleDropdown = () => (showAccountDropdown.value = !showAccountDropdown.value);
const applyFilter = (filterType) => {
  console.log(`Filter applied: ${filterType}`);
  showFilter.value = false;
};
const searchRecipes = () => console.log('Search query:', recipeStore.searchQuery);
const logout = () => {
  authStore.setUser(null);
  authStore.isAuthenticated = false;
  localStorage.removeItem('token');
  router.push('/login');
};

const isLoginOrSignupPage = computed(() =>
  ['/login', '/signup', '/recipeDetail'].includes(route.path)
);
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
/* Ensure the wrapper takes full screen height */
.flex {
  display: flex;
}
.flex-col {
  flex-direction: column;
}
.min-h-screen {
  min-height: 100vh; /* Full viewport height */
}
.flex-grow {
  flex: 1; /* Allow main content to expand */
}
</style>
