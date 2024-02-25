<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import HillControl from "@/components/HillControl.vue";
import { useStationStore } from "@/stores/station";


const { t } = useI18n();
const stationStore = useStationStore()

const hills = computed(() => stationStore.hills)
</script>

<template>
  <div class="flex flex-col" v-for="hill in hills">
    <HillControl :hill="hill" />
  </div>
</template>

