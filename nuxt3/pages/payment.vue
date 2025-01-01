<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center text-blue-600">Make a Payment</h2>
      <form @submit.prevent="submitPayment" class="space-y-4">
        <input
          v-model="email"
          type="email"
          placeholder="Email"
          required
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <input
          v-model="fname"
          type="text"
          placeholder="First Name"
          required
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <input
          v-model="lname"
          type="text"
          placeholder="Last Name"
          required
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <input
          v-model="amount"
          type="number"
          placeholder="Amount"
          required
          min="1"
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />
        <button
          type="submit"
          :disabled="loading"
          class="w-full py-2 bg-blue-600 text-white font-bold rounded-lg hover:bg-blue-700 disabled:opacity-50"
        >
          {{ loading ? 'Processing...' : 'Pay Now' }}
        </button>
      </form>

      <div v-if="loading" class="mt-4 text-center text-gray-600">Processing payment...</div>

      <div v-if="paymentSuccess" class="mt-6 text-center">
        <h3 class="text-xl font-bold text-green-600">Payment Successful!</h3>
        <p class="text-gray-700 mt-2">Thank you for your payment. Transaction ID: {{ paymentSuccess.tx_id }}</p>
        <button
          @click="redirectToHome"
          class="mt-4 py-2 px-4 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          Go to Home
        </button>
      </div>

      <div v-if="paymentError" class="mt-4 text-red-600 text-center">{{ paymentError }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";

const email = ref("");
const fname = ref("");
const lname = ref("");
const amount = ref(0);
const loading = ref(false);
const paymentSuccess = ref(null);
const paymentError = ref("");

const submitPayment = async () => {
  if (!email.value || !fname.value || !lname.value || amount.value <= 0) {
    alert("Please fill in all required fields correctly.");
    return;
  }

  loading.value = true;
  paymentError.value = "";
  paymentSuccess.value = null;

  try {
    const response = await axios.post("http://127.0.0.1:8000/pay", {
      email: email.value,
      fname: fname.value,
      lname: lname.value,
      amount: amount.value,
      rdurl: "http://localhost:3000/paymentComplete",
    });

    const paymentData = response.data;
    const checkoutUrl = paymentData.detail.data.checkout_url;

    if (checkoutUrl) {
      window.location.replace(checkoutUrl); // Redirect to Chapa checkout
      setTimeout(() => {
        redirectToHome();
      }, 60000); // Redirect back to home after 1 minute
    } else {
      paymentError.value = "Checkout URL not received.";
    }
  } catch (error) {
    console.error("Payment initialization failed:", error);
    paymentError.value =
      error.response?.data?.message || "An error occurred while initializing payment.";
  } finally {
    loading.value = false;
  }
};

const redirectToHome = () => {
  window.location.href = "http://localhost:3000/"; // Redirect to home page
};
</script>

<style>
body {
  font-family: 'Inter', sans-serif;
}
</style>
