import { defineStore } from "pinia";
import type { station } from '../../wailsjs/go/models';
import { GetStationHash, LoadStationFile } from '../../wailsjs/go/main/App'

export type StationDef = station.Definition
export type HillDef = station.Hill

export const useStationStore = defineStore('stations', {
  state: () => ({
    hills: [] as station.Hill[],
    hash: ''
  }),
  actions: {
    async loadStation() {
      const def = await LoadStationFile()
      const hash = await GetStationHash()

      this.hills = def.hills
      this.hash = hash
    }
  }
})