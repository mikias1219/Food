<template>
  <div>
    <h2>Transaction Complete</h2>
    <p v-if="loading">Verifying payment...</p>
    <div v-else-if="paymentDetails">
      <h3>Payment Successful</h3>
      <p>Transaction ID: {{ paymentDetails.tx_id }}</p>
      <button @click="redirectToHome">Go to Home</button>
    </div>
    <div v-else-if="error" class="error">
      <p>{{ error }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";

const loading = ref(false);
const paymentDetails = ref(null);
const error = ref("");

const fetchPaymentDetails = async (txRef) => {
  loading.value = true;
  error.value = "";
  try {
    const response = await axios.get(`http://127.0.0.1:8000/payment-complete?tx_ref=${txRef}`);
    paymentDetails.value = response.data;
  } catch (err) {
    error.value = err.response?.data?.detail || "An error occurred.";
  } finally {
    loading.value = false;
  }
};

const redirectToHome = () => {
  window.location.href = "http://localhost:3000/"; // Redirect to home page
};

// Call fetchPaymentDetails(txRef) with the actual transaction reference
fetchPaymentDetails("your-transaction-reference-here");
</script>
