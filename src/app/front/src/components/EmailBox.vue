<template>
  <div
    v-if="email"
    class="rounded-lg border bg-card text-card-foreground shadow-sm w-11/12"
    data-v0-t="card"
  >
    <div class="p-6">
      <div class="flex items-start gap-4">
        <span class="relative flex shrink-0 overflow-hidden rounded-full h-12 w-12">
          <span class="flex h-full w-full items-center justify-center rounded-full bg-muted">
            {{ getInitials(email.From) }}
          </span>
        </span>
        <div class="flex-1">
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <div class="font-medium">{{ email.From }}</div>
              <div class="text-sm text-muted-foreground">{{ email.From }}</div>
            </div>
            <div class="text-sm text-muted-foreground">{{ formatDate(email.Date) }}</div>
          </div>
          <div
            data-orientation="horizontal"
            role="none"
            class="shrink-0 bg-border h-[1px] w-full my-4"
          ></div>
          <div class="prose prose-sm prose-p:leading-normal max-h-[400px] overflow-y-auto">
            <h3 class="mb-4">{{ email.Subject }}</h3>
            <p>{{ email.Body }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="rounded-lg border bg-card text-card-foreground shadow-sm w-11/12 p-6">
    <p>Select an email to view details</p>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import type { Email } from '../models/Email'

defineProps<{
  email: Email | null
}>()

const getInitials = (name: string) => {
  return name
    .split(' ')
    .map((word) => word[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    hour12: true
  })
}
</script>

<style scoped>
/* ... (keep your existing styles) ... */
</style>
