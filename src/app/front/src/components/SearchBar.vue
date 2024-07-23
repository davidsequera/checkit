<template>
  <div class="relative">
    <MagnifyingGlass
      class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground"
    />
    <input
      type="text"
      placeholder="Search emails..."
      v-model="localSearch"
      class="pl-10 pr-4 py-2 rounded-md border border-muted focus:outline-none focus:ring-2 focus:ring-primary focus:border-primary"
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
