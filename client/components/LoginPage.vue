<template>
  <div
    class="min-h-screen bg-gray-100 flex items-center justify-center px-4 py-8"
  >
    <div class="max-w-md w-full bg-white shadow-lg rounded-xl overflow-hidden">
      <div class="bg-blue-600 text-white text-center py-6">
        <Hospital class="mx-auto mb-4" :size="48" />
        <h1 class="text-2xl font-bold">Medical Portal Login</h1>
      </div>

      <form @submit.prevent="handleLogin" class="p-6 space-y-6">
        <!-- Role Selection -->
        <!-- <div class="flex justify-center space-x-4 mb-6">
          <button
            v-for="userRole in ['receptionist', 'doctor']"
            :key="userRole"
            type="button"
            @click="role = userRole"
            :class="[
              'px-4 py-2 rounded-full text-sm uppercase font-semibold transition-colors',
              role === userRole
                ? 'bg-blue-600 text-white'
                : 'bg-gray-200 text-gray-700 hover:bg-gray-300',
            ]"
          >
            {{ userRole }}
          </button>
        </div> -->

        <!-- Username Input -->
        <div class="relative">
          <div
            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
          >
            <User class="text-gray-400" :size="20" />
          </div>
          <input
            type="text"
            placeholder="Username"
            v-model="userName"
            required
            class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <!-- Password Input -->
        <div class="relative">
          <div
            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
          >
            <LockKeyhole class="text-gray-400" :size="20" />
          </div>
          <input
            type="password"
            placeholder="Password"
            v-model="password"
            required
            class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <!-- Login Button -->
        <button
          type="submit"
          class="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-semibold uppercase tracking-wider"
        >
          Login
        </button>

        <!-- Forgot Password -->
        <div class="text-center">
          <a href="#" class="text-sm text-blue-600 hover:underline">
            Forgot Password?
          </a>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Hospital, User, LockKeyhole } from "lucide-vue-next";

const userName = ref("");
const password = ref("");
const role = ref("receptionist");

const handleLogin = async () => {
  const API_URL = useRuntimeConfig().public.API_BASE_URL;
  // console.log("API_URL:", API_URL);
  try {
    const response = await fetch(`${API_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Username: userName.value,
        Password: password.value,
        // role: role.value,
      }),
    });

    const data = await response.json();
    console.log("Data:", data);

    if (response.status === 200) {
      console.log("Login successful");
      sessionStorage.setItem("token", data.token);
      sessionStorage.setItem("role", data.role);
      sessionStorage.setItem("username", data.name);
      navigateTo("/dashboard");
    } else {
      // Login failed
      console.error("Login failed");
    }
    console.log("Login response:", response);
  } catch (error) {
    console.error("Login failed:", error);
  }
  console.log("Login attempt:", {
    Username: userName.value,
    Password: password.value,
    // role: role.value,
  });
};
</script>
