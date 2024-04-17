<script lang="ts" setup>
// import { InitApp } from '~/utils/initApp'
// InitApp()
const locale = useState<string>('locale.i18n')
</script>

<template>
  <Html :lang="locale">
  <Body>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  </Body>
  </Html>
</template>

<style lang="postcss">
body {
  @apply bg-gray-50 dark:bg-gray-800;
}
.global-text {
  @apply text-gray-900 dark:text-gray-50;
}
</style>