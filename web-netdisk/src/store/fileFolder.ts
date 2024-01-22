import {defineStore} from "pinia";
import {reactive} from "vue";

export const useFileFolderStore = defineStore('file-folder', () => {
    const selectedItems: { files: number[], folders: number[] } = reactive({
        files: [],
        folders: []
    })

    function selectChange(ids: number[], forFile: boolean) {
        if (forFile) {
            selectedItems.files = ids
        } else {
            selectedItems.folders = ids
        }
    }

    return {
        selectedItems,
        selectChange
    }
})