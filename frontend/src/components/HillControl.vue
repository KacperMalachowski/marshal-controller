<script setup lang="ts">
import type { HillDef } from "@/stores/station";
import { defineProps, ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { SetSignal } from "../../wailsjs/go/main/App";

type Props = {
  hill: HillDef;
};

const { hill } = defineProps<Props>();
const setState = (signal: string) => SetSignal(hill, signal);

const signal = ref('Off');

EventsOn("message", (message: string) => {
  if (!message.includes(`${hill.signal}:`)) return
  const splitted = message.split(':')

  if (splitted.length !== 2) return
  signal.value = splitted[1]

});
</script>

<template>
  <div class="w-100 rounded overflow-hidden shadow-md mb-5">
    <div class="px-6 py-4">
      <div class="font-bold text-xl mb-2 border-b-2">
        Main signal: {{ hill.signal }}
      </div>
    </div>
    <div class="px-6 pt-4 pb-2">
      <button
        @click="setState('Off')"
        class="inline-block  rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Off' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt0 - Off
      </button>
      <button
        @click="setState('Rt1')"
        class="inline-block rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Rt1' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt1 - Push forbidden
      </button>
      <button
        @click="setState('Rt2')"
        class="inline-block rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Rt2' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt2 - Push slowly
      </button>

      <button
        @click="setState('Rt3')"
        class="inline-block rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Rt3' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt3 - Push
      </button><button
        @click="setState('Rt4')"
        class="inline-block rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Rt4' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt4 - Reverse
      </button><button
        @click="setState('Rt5')"
        class="inline-block rounded-full px-3 py-1 font-semibold mr-2 mb-2"
        :class="signal == 'Rt5' ? 'bg-green-300' : 'text-gray-700 bg-gray-200'"
      >
        Rt5 - Push towards hill
      </button>
    </div>
  </div>
</template>
