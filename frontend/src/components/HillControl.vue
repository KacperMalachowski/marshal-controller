<script setup lang="ts">
import type { HillDef } from '@/stores/station';
import { defineProps } from 'vue';
import { SetSignal } from '../../wailsjs/go/main/App';


type Props = {
  hill: HillDef
}

const { hill } = defineProps<Props>()
const setState = (signal: string) => SetSignal(hill, signal)
</script>

<template>
<ul class="flex">
  <li class="flex-1 m-2 p-6 justify-between border-solid">
    <button @click="setState('Off')">Rt0 - Off</button>
    <button @click="setState('Rt1')">Rt1 - Push forbidden</button>
    <button @click="setState('Rt2')">Rt2 - Push slowly</button>
    <button @click="setState('Rt3')">Rt3 - Push</button>
    <button @click="setState('Rt4')">Rt4 - Reverse</button>
    <button @click="setState('Rt5')">Rt5 - Move towards hill</button>
  </li>
</ul>
</template>