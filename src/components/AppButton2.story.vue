<script lang="ts" setup>
import {ref} from 'vue'
import {QBtn} from 'quasar'

const isDarkMode = ref(false)

const handleClick = () => {
  alert('Button clicked!')
}

</script>

<template>
  <Story title="Button">
    <template #default>
      <div>
        <!-- Basic Usage -->
        <h2>Basic Usage</h2>
        <q-btn label="Default Button"/>

        <!-- Variants -->
        <h2>Variants</h2>
        <q-btn color="primary" label="Primary"/>
        <q-btn color="secondary" label="Secondary"/>

        <!-- Sizes -->
        <h2>Sizes</h2>
        <q-btn size="small" label="Small Button"/>
        <q-btn size="large" label="Large Button"/>

        <!-- Icons -->
        <h2>Icons</h2>
        <q-btn icon="favorite" label="Icon Button"/>

        <!-- Loading State -->
        <h2>Loading State</h2>
        <q-btn :loading="true" label="Loading..."/>

        <!-- Disabled State -->
        <h2>Disabled State</h2>
        <q-btn :disable="true" label="Disabled Button"/>

        <!-- Button Groups -->
        <h2>Button Groups</h2>
        <q-btn-group>
          <q-btn label="One"/>
          <q-btn label="Two"/>
          <q-btn label="Three"/>
        </q-btn-group>

        <!-- Custom Styles -->
        <h2>Custom Styles</h2>
        <q-btn class="custom-button" label="Custom Styled Button"/>

        <!-- Events and Actions -->
        <h2>Events and Actions</h2>
        <q-btn @click="handleClick" label="Click Me"/>

        <!-- Accessibility -->
        <h2>Accessibility</h2>
        <q-btn aria-label="Accessible Button" label="Accessible Button"/>

        <!-- Theming -->
        <h2>Theming</h2>
        <div :class="{ 'dark-mode': isDarkMode }">
          <q-btn color="primary" label="Themed Button"/>
        </div>
      </div>
    </template>
  </Story>
</template>

<style>
.custom-button {
  background-color: #42b983;
  color: white;
}

.dark-mode {
  background-color: #333;
  color: white;
}
</style>