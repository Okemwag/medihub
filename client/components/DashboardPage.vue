<script setup>
import { ref, computed, reactive } from "vue";
import {
  Users,
  Calendar,
  FilePlus,
  FileText,
  Stethoscope,
  UserPlus,
  LogOut,
  Menu,
  X,
} from "lucide-vue-next";
// import PatientListModal from "./patientListModal.vue";

const userName = ref("John Doe");
const activeRole = ref(1);
const sidebarOpen = ref(false);

const displayUsername = sessionStorage.getItem("username");
const userRole = sessionStorage.getItem("role");

// console.log("userRole:", typeof(userRole), userRole);

userName.value = displayUsername;
activeRole.value = Number(userRole) || userRole;

// console.log("activeRole:", typeof activeRole.value, activeRole.value);

const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value;
};

const setActiveRole = (role) => {
  activeRole.value = role;
  sidebarOpen.value = false; // Close sidebar on mobile after role selection
};

const receptionistMenuItems = [
  { icon: UserPlus, label: "Register Patient" },
  { icon: Users, label: "Patient List" },
  { icon: Calendar, label: "Appointments" },
];

const doctorMenuItems = [
  { icon: FileText, label: "Patient Records", route: "/patient-records" },
  { icon: Stethoscope, label: "Consultations", route: "/consultations" },
  { icon: Calendar, label: "Schedule", route: "/doctor-schedule" },
];

const menuItems = computed(() =>
  activeRole.value === 2 ? receptionistMenuItems : doctorMenuItems
);

const dashboardCards = computed(() => {
  if (activeRole.value === 2) {
    return [
      {
        icon: UserPlus,
        iconClass: "text-blue-600",
        title: "New Patient Registration",
        description: "Total Registrations: 42",
      },
      {
        icon: Calendar,
        iconClass: "text-green-600",
        title: "Today's Appointments",
        description: "Scheduled: 15",
      },
      {
        icon: Users,
        iconClass: "text-purple-600",
        title: "Patient Management",
        description: "Active Patients: 256",
      },
    ];
  } else {
    return [
      {
        icon: Stethoscope,
        iconClass: "text-blue-600",
        title: "Today's Consultations",
        description: "Pending: 8",
      },
      {
        icon: FileText,
        iconClass: "text-green-600",
        title: "Patient Records",
        description: "Total Records: 342",
      },
      {
        icon: Calendar,
        iconClass: "text-purple-600",
        title: "Upcoming Schedule",
        description: "Appointments: 12",
      },
    ];
  }
});

const modals = reactive({
  registerPatient: false,
  appointments: false,
  patientsList: false,
});

const handleMenuAction = (label) => {
  switch (label) {
    case "Register Patient":
      modals.registerPatient = true;
      break;
    case "Appointments":
      modals.appointments = true;
      break;
    case "Patient List":
      modals.patientsList = true;
      break;
  }
};

// const submitPatientRegistration = async () => {

// };

const closeModal = (modalName) => {
  modals[modalName] = false;
};
</script>
<template>
  <div class="flex flex-col md:flex-row h-screen bg-gray-100">
    <!-- Mobile Header -->
    <div
      class="md:hidden bg-white shadow-md p-4 flex justify-between items-center"
    >
      <div class="flex items-center">
        <!-- <img 
          src="https://pixabay.com/photos/doctor-gray-hair-experience-doctor-2337835/" 
          alt="User Profile" 
          class="w-10 h-10 rounded-full mr-3"
        /> -->
        <div>
          <h2 class="font-semibold">{{ userName }}</h2>
          <p class="text-xs text-gray-500">
            <!-- {{ activeRole.charAt(0).toUpperCase() + activeRole.slice(1) }} -->
            {{ activeRole === 2 ? "Receptionist" : "Doctor" }}
          </p>
        </div>
      </div>
      <button @click="toggleSidebar" class="focus:outline-none">
        <Menu :size="24" />
      </button>
    </div>

    <!-- Sidebar -->
    <div
      :class="[
        'fixed inset-y-0 left-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out',
        'md:relative md:translate-x-0',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full',
      ]"
    >
      <!-- Close Button for Mobile -->
      <button
        @click="toggleSidebar"
        class="md:hidden absolute top-4 right-4 focus:outline-none"
      >
        <X :size="24" />
      </button>

      <!-- Desktop User Profile -->
      <div class="hidden md:flex items-center p-6">
        <img
          src="https://pixabay.com/photos/doctor-gray-hair-experience-doctor-2337835/"
          alt="User Profile"
          class="w-12 h-12 rounded-full mr-4"
        />
        <div>
          <h2 class="font-semibold">{{ userName }}</h2>
          <p class="text-sm text-gray-500">
            <!-- {{ activeRole.charAt(0).toUpperCase() + activeRole.slice(1) }} -->
            {{ activeRole === 2 ? "Receptionist" : "Doctor" }}
          </p>
        </div>
      </div>

      <!-- Role Toggle -->
      <!-- <div class="flex p-4 justify-center">
        <button
          v-for="role in ['receptionist', 'doctor']"
          :key="role"
          @click="setActiveRole(role)"
          :class="[
            'px-3 py-1 rounded-full text-sm mr-2 transition-colors',
            activeRole === role
              ? 'bg-blue-600 text-white'
              : 'bg-gray-200 text-gray-700 hover:bg-gray-300',
          ]"
        >
          {{ role.charAt(0).toUpperCase() + role.slice(1) }}
        </button>
      </div> -->

      <!-- Menu Items  -->
      <nav class="p-2">
        <button
          v-for="item in menuItems"
          :key="item.label"
          @click="handleMenuAction(item.label)"
          class="w-full flex items-center p-3 hover:bg-gray-100 rounded-lg text-left cursor-pointer"
        >
          <component :is="item.icon" class="mr-3" :size="20" />
          <span class="text-sm">{{ item.label }}</span>
        </button>

        <button
          @click="logout"
          class="w-full flex items-center p-3 hover:bg-gray-100 rounded-lg text-red-600"
        >
          <LogOut class="mr-3" :size="20" />
          <span>Logout</span>
        </button>
      </nav>
    </div>

    <RegisterPatientModal
      :is-open="modals.registerPatient"
      @close="closeModal('registerPatient')"
    />

    <PatientListModal
      :is-open="modals.patientsList"
      @close="closeModal('patientsList')"
    />

    <!-- Main Content -->
    <div class="flex-1 p-4 md:p-8 overflow-y-auto">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="card in dashboardCards"
          :key="card.title"
          class="bg-white p-4 md:p-6 rounded-lg shadow"
        >
          <component
            :is="card.icon"
            :class="card.iconClass"
            :size="32"
            md:size="40"
          />
          <h3 class="text-lg md:text-xl font-semibold mt-2 md:mt-4">
            {{ card.title }}
          </h3>
          <p class="text-sm">{{ card.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
