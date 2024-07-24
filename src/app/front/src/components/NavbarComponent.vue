<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header class="flex h-16 w-full items-center justify-between bg-background px-4 md:px-6">
    <a class="flex items-center gap-2" href="#" rel="ugc">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="h-6 w-6"
      >
        <path d="m8 3 4 8 5-5 5 15H2L8 3z"></path>
      </svg>
      <span class="text-lg font-semibold">Acme Inc</span>
    </a>
    <nav class="hidden items-center gap-6 md:flex">
      <RouterLink
        class="font-medium hover:text-sky-700 transition-colors hover:text-primary"
        to="/"
      >
        Home
      </RouterLink>
      <RouterLink
        class="font-medium hover:text-sky-700 transition-colors hover:text-primary"
        to="/about"
      >
        About
      </RouterLink>
      <RouterLink
        class="font-medium hover:text-sky-700 transition-colors hover:text-primary"
        to="/"
      >
        Services
      </RouterLink>
      <RouterLink
        class="font-medium hover:text-sky-700 transition-colors hover:text-primary"
        to="/"
      >
        Contact
      </RouterLink>
    </nav>
    <button
      class="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 w-10 md:hidden"
      type="button"
      aria-haspopup="dialog"
      aria-expanded="false"
      aria-controls="radix-:r6:"
      data-state="closed"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="h-6 w-6"
      >
        <line x1="4" x2="20" y1="12" y2="12"></line>
        <line x1="4" x2="20" y1="6" y2="6"></line>
        <line x1="4" x2="20" y1="18" y2="18"></line>
      </svg>
      <span class="sr-only">Toggle navigation</span>
    </button>
  </header>
  <template />
</template>
