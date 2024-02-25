import { defineStore } from "pinia";
import type { station } from '../../wailsjs/go/models';
import { LoadStationFile } from '../../wailsjs/go/main/App'

export type StationDef = station.Definition
export type HillDef = station.Hill

export const useStationStore = defineStore('stations', {
  state: () => ({
    hills: [] as station.Hill[] 
  }),
  actions: {
    async loadStation() {
      const def = await LoadStationFile()

      console.log(def)
      this.hills = def.hills
    }
  }
})