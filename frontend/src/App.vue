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
          class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0"
        >
          {{ connected ? 'Disconnect' : 'Connect' }}
        </button>
        <button
          @click="() => stationStore.loadStation()"
          class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0"
        >
          Load station definition
        </button>
      </div>
    </div>
  </nav>
  <RouterView />
</template>
