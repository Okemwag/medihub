<script setup>
import { ref, reactive } from "vue";
import { X } from "lucide-vue-next";

const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["close"]);

const patient = reactive({
  FirstName: "",
  LastName: "",
  Email: "",
  ContactNumber: "",
  DateOfBirth: "",
  Gender: "",
  Address: "",
  MedicalHistory: "",
});

const closeModal = () => {
  emit("close");
};

const submitRegistration = async () => {
  const authToken = sessionStorage.getItem("token");
  console.log("authToken:", authToken);
  try {
    // Validate patient data
    if (!patient.FirstName || !patient.LastName) {
      alert("First and last name are required");
    }

    if (!patient.Email) {
      alert("Email is required");
    }

    if (!patient.ContactNumber) {
      alert("Your contact is required");
    }

    if (!patient.Address) {
      alert("Your address is required");
    }

    if (!patient.MedicalHistory) {
      alert("Your medical history is required");
    }

    if (!patient.DateOfBirth) {
      alert("Your date of birth is required");
    }

    if (!patient.Gender) {
      alert("Your gender is required");
    }

    const API_URL = useRuntimeConfig().public.API_BASE_URL;
    // console.log("API_URL:", API_URL);
    try {
      const response = await fetch(`${API_URL}/patients`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${authToken}`,
        },
        body: JSON.stringify({
          patient,
        }),
      });

      const data = response.json();
      console.log("data:", data);
      if (response.status === 200) {
        // Registration successful
        console.log("Patient registered successfully");
      } else {
        // Registration failed
        console.error("Patient Registration failed");
      }
      console.log("Registration response:", response);
    } catch (error) {
      console.error("Patient Registration failed:", error);
    }
    console.log("Registration attempt:", patient);
    // Emit patient data to parent component
    // emit("submit", { ...patient });

    // Reset form and close modal
    Object.keys(patient).forEach((key) => (patient[key] = ""));
    closeModal();
  } catch (error) {
    console.error("Registration error:", error);
  }
};
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50"
  >
    <div class="bg-white rounded-lg shadow-xl w-full max-w-md p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-xl font-semibold">Patient Registration</h2>
        <button @click="closeModal" class="text-gray-500 hover:text-gray-700">
          <X :size="24" />
        </button>
      </div>

      <form @submit.prevent="submitRegistration">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >First Name</label
            >
            <input
              v-model="patient.FirstName"
              type="text"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >Last Name</label
            >
            <input
              v-model="patient.LastName"
              type="text"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
        </div>

        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700">Email</label>
          <input
            v-model="patient.Email"
            type="email"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
          />
        </div>

        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700">Phone</label>
          <input
            v-model="patient.ContactNumber"
            type="tel"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
          />
        </div>

        <div class="grid grid-cols-2 gap-4 mt-4">
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >Date of Birth</label
            >
            <input
              v-model="patient.DateOfBirth"
              type="date"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >Gender</label
            >
            <select
              v-model="patient.Gender"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            >
              <option value="">Select</option>
              <option value="male">Male</option>
              <option value="female">Female</option>
              <option value="other">Other</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4 mt-4">
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >Address</label
            >
            <input
              v-model="patient.Address"
              type="text"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700"
              >Medical History</label
            >
            <input
              v-model="patient.MedicalHistory"
              type="text"
              required
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
            />
          </div>
        </div>

        <div class="mt-6 flex justify-end space-x-3">
          <button
            type="button"
            @click="closeModal"
            class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            Register Patient
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
