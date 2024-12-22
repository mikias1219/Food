<template>
    <div>
      <h2>Verify Transaction</h2>
      <form @submit.prevent="submitVerification">
        <input v-model="txRef" type="text" placeholder="Transaction Reference" required />
        <button type="submit">Verify</button>
      </form>
  
      <div v-if="verificationResponse">
        <h3>Verification Status</h3>
        <p>Status: {{ verificationResponse.status }}</p>
        <p>Message: {{ verificationResponse.message }}</p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import axios from 'axios'
  
  const txRef = ref('')
  const verificationResponse = ref(null)
  
  const submitVerification = async () => {
    try {
      const response = await axios.post('http://127.0.0.1:8000/verify', {
        ref_num: txRef.value
      })
      verificationResponse.value = response.data
    } catch (error) {
      console.error('Verification failed:', error)
    }
  }
  </script>
  