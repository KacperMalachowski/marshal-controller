<script setup lang="ts">
import { ref } from "vue";
import { Connect, Disconnect } from "../wailsjs/go/main/App";
import { useStationStore } from "./stores/station";

const connected = ref(false);
const stationStore = useStationStore()

function handleConnect() {
  if (connected.value) {
    connected.value = false
    Disconnect()
  } else {
    Connect("127.0.0.1:7424")
    connected.value = true
  }
}
</script>

<template>
  <nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6">
    <div class="flex items-center flex-shrink-0 text-white mr-6">
      <span class="font-semibold text-xl tracking-tight"
        >Marshaller Controller</span
      >
    </div>
    <div class="flex-grow flex items-center w-auto">
      <div class="text-sm flex-grow"></div>
      <div>
        <button
          @click="handleConnect"
          :disabled="stationStore.hills.length === 0"
          class="btn btn-main"
        >
          {{ connected ? 'Disconnect' : 'Connect' }}
        </button>
        <button
          @click="() => stationStore.loadStation()"
          class="btn btn-secondary"
        >
          Load station definition
        </button>
      </div>
    </div>
  </nav>
  <RouterView />
</template>

<style lang="scss">
  .btn {
    @apply font-bold py-2 px-4 rounded mr-2 ml-2;
  }
  .btn-main {
    @apply bg-gray-300;
  }
  .btn-main:hover {
    @apply bg-gray-700 text-white;
  }
  .btn-secondary {
    @apply bg-transparent border border-gray-700 hover:border-transparent;
  }
  .btn-secondary:hover {
    @apply bg-gray-300;
  }
</style>
