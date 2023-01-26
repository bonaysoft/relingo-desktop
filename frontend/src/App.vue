<template>
  <v-app>
    <v-navigation-drawer permanent>
      <v-list-item :prepend-avatar="uInfo.avatar" :title="uInfo.email" nav>
        <template v-slot:append>
          <v-btn variant="text" icon="mdi-chevron-left"></v-btn>
        </template>
      </v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-home-city" title="我的生词" value="home" to="home"></v-list-item>
        <v-list-item prepend-icon="mdi-account" title="全部单词" value="account" to="words"></v-list-item>
        <v-list-item prepend-icon="mdi-account-group-outline" title="已掌握单词" value="users" to="mastered"></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { GetUserInfo } from '@/../wailsjs/go/relingo/Client';

const uInfo = ref<any>({})
onMounted(async () => {
  uInfo.value = await GetUserInfo()
})


</script>
