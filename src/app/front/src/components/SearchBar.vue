<template>
  <div class="relative">
    <MagnifyingGlass
      class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground"
    />
    <input
      type="text"
      placeholder="Search emails..."
      v-model="localSearch"
      class="transition-colors duration-300 ease-in-out focus-visible:border-sky-800 w-full outline-none rounded-full pl-10 pr-4 py-2 border"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue'

import MagnifyingGlass from '@heroicons/vue/24/outline/MagnifyingGlassIcon'

export default defineComponent({
  name: 'SearchBar',
  components: {
    MagnifyingGlass
  },
  props: {
    modelValue: String
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const localSearch = ref(props.modelValue)

    watch(localSearch, (newValue) => {
      emit('update:modelValue', newValue)
    })

    return {
      localSearch
    }
  }
})
</script>
