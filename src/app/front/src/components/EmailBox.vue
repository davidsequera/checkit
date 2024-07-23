<template>
  <div
    class="rounded-lg border bg-card text-card-foreground shadow-sm w-full max-w-2xl"
    data-v0-t="card"
  >
    <div class="p-6">
      <div class="flex items-start gap-4">
        <span class="relative flex shrink-0 overflow-hidden rounded-full h-12 w-12"
          ><span class="flex h-full w-full items-center justify-center rounded-full bg-muted"
            >OD</span
          ></span
        >
        <div class="flex-1">
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <div class="font-medium">Olivia Davis</div>
              <div class="text-sm text-muted-foreground">olivia.davis@vercel.com</div>
            </div>
            <div class="text-sm text-muted-foreground">Oct 08, 2023 9:15 AM</div>
          </div>
          <div
            data-orientation="horizontal"
            role="none"
            class="shrink-0 bg-border h-[1px] w-full my-4"
          ></div>
          <div class="prose prose-sm prose-p:leading-normal max-h-[400px] overflow-y-auto">
            <h3 class="mb-4">Question about Budget</h3>
            <p>
              Hi, let's have a meeting tomorrow to discuss the project. I've been reviewing the
              project details and have some ideas I'd like to share. It's crucial that we align on
              our next steps to ensure the project's success.
            </p>
            <p>
              Please come prepared with any questions or insights you may have. Looking forward to
              our meeting!
            </p>
            <p>Best,<br />Olivia</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EmailCard'
}
</script>

<style scoped>
.card {
  border: solid 1px blue;
  border-radius: 1rem;
  /* Add your card styles here */
}

.card-content {
  border: solid 1px blue;
  border-radius: 1rem;
  /* Add your card content styles here */
}

.avatar {
  border: solid 1px blue;
  border-radius: 1rem;

  /* Add your avatar styles here */
}

.avatar-image {
  border: solid 1px blue;
  border-radius: 1rem;

  /* Add your avatar image styles here */
}

.avatar-fallback {
  border: single 1px blue;
  border-radius: 1rem;

  /* Add your avatar fallback styles here */
}

.separator {
  border: solid 1px blue;
  border-radius: 1rem;

  /* Add your separator styles here */
}
</style>
