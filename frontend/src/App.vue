<template>
  <v-app v-if="!ready">
    <Login />
  </v-app>
  <v-app v-else>
    <v-navigation-drawer permanent>
      <v-list-item :prepend-avatar="uInfo.avatar" :title="uInfo.email" nav>
        <template v-slot:append>
          <v-btn variant="text" icon="mdi-chevron-left"></v-btn>
        </template>
      </v-list-item>

      <v-divider></v-divider>

      <!-- <v-list density="compact" nav>
      </v-list> -->

      <v-list>

        <v-list-group value="Users">
          <template v-slot:activator="{ props }">
            <v-list-item v-bind="props" prepend-icon="mdi-home-city" title="我的生词" value="home"></v-list-item>
          </template>
          <v-list-item title="今日生词" value="today" :to="{ name: 'home', query: { type: 'today' } }"></v-list-item>
          <v-list-item title="昨日生词" value="yesterday" :to="{ name: 'home', query: { type: 'yesterday' } }"></v-list-item>
          <v-list-item title="本周生词" value="weekly" :to="{ name: 'home', query: { type: 'weekly' } }"></v-list-item>
          <v-list-item title="全部生词" value="all" :to="{ name: 'home', query: { type: 'all' } }"></v-list-item>
        </v-list-group>
        <v-list-group value="NewWords">
          <template v-slot:activator="{ props }">
            <v-list-item v-bind="props" prepend-icon="mdi-home-city" title="全部单词" value="words" to="words"></v-list-item>
          </template>

          <!-- <v-list-item title="今日生词" value="today" to="words"></v-list-item>
          <v-list-item title="昨日生词" value="yesterday" to="mastered"></v-list-item>
          <v-list-item title="本周生词" value="weekly" to="mastered"></v-list-item>
          <v-list-item title="全部生词" value="all" to="mastered"></v-list-item> -->
        </v-list-group>
        <v-list-group value="MasteredWords">
          <template v-slot:activator="{ props }">
            <v-list-item v-bind="props" prepend-icon="mdi-home-city" title="已掌握单词" value="mastered"
              to="mastered"></v-list-item>
          </template>

          <v-list-item title="最近掌握" value="account" to="mastered"></v-list-item>
          <v-list-item title="需要复习" value="users" to="mastered"></v-list-item>
        </v-list-group>
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
import Login from "@/components/Login.vue";
import { ref, onMounted } from 'vue';
import { Ready, GetUserInfo } from '@/../wailsjs/go/relingo/Client';

const ready = ref(false)
const uInfo = ref<any>({})
onMounted(async () => {
  ready.value = await Ready()
  if (!ready.value) {
    return
  }

  uInfo.value = await GetUserInfo()
})
</script>
