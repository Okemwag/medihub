<script setup>
import { ref, computed } from "vue";
import {
  Search,
  ChevronDown,
  ChevronUp,
  Edit,
  Trash,
  X,
} from "lucide-vue-next";

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true,
  },
});

defineEmits(["close"]);

// Table headers configuration
const tableHeaders = [
  { key: "id", label: "ID" },
  { key: "name", label: "Patient Name" },
  { key: "phone", label: "Phone" },
  { key: "dateOfBirth", label: "Date of Birth" },
  { key: "status", label: "Status" },
];

// State
const searchTerm = ref("");
const filterStatus = ref("all");
const sortColumn = ref("name");
const sortDirection = ref("asc");
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Func to fetch patients
const fetchPatients = async () => {
  const API_URL =
    useRuntimeConfig().public.API_BASE_URL || "http://localhost:8000";
  console.log("API_URL:", API_URL);
  try {
    const response = await fetch(`${API_URL}/patients`);
    const data = await response.json();
    patients.value = data;
  } catch (error) {
    console.error("Failed to fetch patients:", error);
  }
};

// Sample data
const patients = ref([
  {
    id: "001",
    name: "John Smith",
    email: "john.smith@email.com",
    phone: "+1 234-567-8901",
    dateOfBirth: "1985-06-15",
    status: "Active",
  },
]);

// Computed properties
const filteredPatients = computed(() => {
  let filtered = [...patients.value];

  // Apply search
  if (searchTerm.value) {
    const search = searchTerm.value.toLowerCase();
    filtered = filtered.filter(
      (patient) =>
        patient.name.toLowerCase().includes(search) ||
        patient.email.toLowerCase().includes(search) ||
        patient.phone.includes(search)
    );
  }

  // Apply status filter
  if (filterStatus.value !== "all") {
    filtered = filtered.filter(
      (patient) => patient.status.toLowerCase() === filterStatus.value
    );
  }

  // Apply sorting
  filtered.sort((a, b) => {
    const aValue = a[sortColumn.value];
    const bValue = b[sortColumn.value];

    if (sortDirection.value === "asc") {
      return aValue.localeCompare(bValue);
    }
    return bValue.localeCompare(aValue);
  });

  return filtered;
});

// Pagination computed properties
const totalPatients = computed(() => filteredPatients.value.length);
const totalPages = computed(() =>
  Math.ceil(totalPatients.value / itemsPerPage.value)
);
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value);
const endIndex = computed(() =>
  Math.min(startIndex.value + itemsPerPage.value, totalPatients.value)
);
const paginatedPatients = computed(() =>
  filteredPatients.value.slice(startIndex.value, endIndex.value)
);

// Methods
const sortBy = (column) => {
  if (sortColumn.value === column) {
    sortDirection.value = sortDirection.value === "asc" ? "desc" : "asc";
  } else {
    sortColumn.value = column;
    sortDirection.value = "asc";
  }
};

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
};

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
};

const editPatient = (patient) => {
  // Implement edit functionality
  console.log("Edit patient:", patient);
};

const deletePatient = (patient) => {
  // Implement delete functionality
  console.log("Delete patient:", patient);
};
</script>
<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div
      class="fixed inset-0 bg-black bg-opacity-50 transition-opacity"
      @click="$emit('close')"
    ></div>

    <!-- Modal Content -->
    <div class="flex items-center justify-center min-h-screen p-4">
      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-6xl">
        <!-- Header -->
        <div class="flex items-center justify-between p-6 border-b">
          <h2 class="text-2xl font-semibold">Patient List</h2>
          <button
            @click="$emit('close')"
            class="text-gray-500 hover:text-gray-700"
          >
            <X :size="24" />
          </button>
        </div>

        <!-- Search and Filter Section -->
        <div class="p-6 border-b">
          <div class="flex flex-col md:flex-row gap-4">
            <!-- Search -->
            <div class="relative flex-1">
              <input
                v-model="searchTerm"
                type="text"
                placeholder="Search patients..."
                class="w-full pl-10 pr-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <Search
                class="absolute left-3 top-2.5 text-gray-400"
                :size="20"
              />
            </div>
            <!-- Filter -->
            <select
              v-model="filterStatus"
              class="px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="all">All Patients</option>
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
            </select>
          </div>
        </div>

        <!-- Table -->
        <div class="p-6 overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  v-for="header in tableHeaders"
                  :key="header.key"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
                  @click="sortBy(header.key)"
                >
                  <div class="flex items-center gap-2">
                    {{ header.label }}
                    <span v-if="sortColumn === header.key">
                      <ChevronUp v-if="sortDirection === 'asc'" :size="16" />
                      <ChevronDown v-else :size="16" />
                    </span>
                  </div>
                </th>
                <th class="px-6 py-3 text-right">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr
                v-for="patient in filteredPatients"
                :key="patient.id"
                class="hover:bg-gray-50"
              >
                <td class="px-6 py-4 whitespace-nowrap">{{ patient.id }}</td>
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">
                        {{ patient.name }}
                      </div>
                      <div class="text-sm text-gray-500">
                        {{ patient.email }}
                      </div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">{{ patient.phone }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  {{ patient.dateOfBirth }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
                    :class="[
                      'px-2 py-1 text-xs rounded-full',
                      patient.status === 'Active'
                        ? 'bg-green-100 text-green-800'
                        : 'bg-gray-100 text-gray-800',
                    ]"
                  >
                    {{ patient.status }}
                  </span>
                </td>
                <td
                  class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium"
                >
                  <button
                    @click="editPatient(patient)"
                    class="text-blue-600 hover:text-blue-900 mr-3"
                  >
                    <Edit :size="18" />
                  </button>
                  <button
                    @click="deletePatient(patient)"
                    class="text-red-600 hover:text-red-900"
                  >
                    <Trash :size="18" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>

          <!-- Pagination -->
          <div class="flex items-center justify-between mt-6">
            <div class="text-sm text-gray-700">
              Showing {{ startIndex + 1 }} to {{ endIndex }} of
              {{ totalPatients }} results
            </div>
            <div class="flex gap-2">
              <button
                @click="previousPage"
                :disabled="currentPage === 1"
                class="px-4 py-2 border rounded-lg disabled:opacity-50"
              >
                Previous
              </button>
              <button
                @click="nextPage"
                :disabled="currentPage === totalPages"
                class="px-4 py-2 border rounded-lg disabled:opacity-50"
              >
                Next
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
