<script setup lang="ts">
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import HillControl from "@/components/HillControl.vue";
import { useStationStore } from "@/stores/station";


const { t } = useI18n();
const stationStore = useStationStore()

const hills = computed(() => stationStore.hills)
const hash = computed(() => stationStore.hash)
</script>

<template>
  <div v-if="hash !== ''" class="text-center">Station hash {{ hash }}</div>
  <div class="flex flex-col" v-for="hill in hills">
    <HillControl :hill="hill" />
  </div>
</template>

