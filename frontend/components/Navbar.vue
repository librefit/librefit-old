<template>
  <nav>
    <v-app-bar app clipped-left color="primary">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <span class="title ml-3 mr-5"
        >Libre&nbsp;<span class="font-weight-light">Fit</span></span
      >
      <v-spacer></v-spacer>
      <v-spacer></v-spacer>

      <v-menu offset-y>
        <template v-slot:activator="{ attrs, on }">
          <v-btn icon v-bind="attrs" v-on="on">
            <div v-if="LoggedIn">
              <v-avatar v-if="Avatar" color="primary">
                <img :src="'http://localhost:4000/img/' + Avatar" />
              </v-avatar>
              <v-avatar v-else>
                {{ NameAvatar }}
              </v-avatar>
            </div>
          </v-btn>
        </template>

        <v-list>
          <v-list-item to="/user/settings" link>
            <v-list-item-title> Settings </v-list-item-title>
          </v-list-item>
          <v-list-item href="https://github.com/librefit/librefit" target="_blank">
            <v-list-item-title> About LibreFit </v-list-item-title>
          </v-list-item>
          <v-list-item @click="logOut" link>
            <v-list-item-title> Logout </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      app
      clipped
      color="grey lighten-4"
      width="250"
    >
      <v-list>
        <template v-for="(item, i) in genericMenu">
          <v-subheader v-if="item.header" :key="item.header" inset>
            {{ item.header }}
          </v-subheader>
          <v-divider
            v-else-if="item.divider"
            :key="i + 'first'"
            inset
          ></v-divider>

          <v-list-item
            v-if="!item.submenu"
            :key="i + 'second'"
            :to="item.route"
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item>
          <v-list-group
            v-else-if="item.submenu"
            :key="i + 'third'"
            :prepend-icon="item.icon"
            no-action
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>{{ item.text }}</v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item
              v-for="(subitem, s) in item.submenu"
              :key="s"
              :to="subitem.route"
            >
              <v-list-item-title v-text="subitem.text"></v-list-item-title>
              <v-list-item-icon>
                <v-icon v-text="subitem.icon"></v-icon>
              </v-list-item-icon>
            </v-list-item>
          </v-list-group>
        </template>
      </v-list>
    </v-navigation-drawer>
  </nav>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    drawer: false,
    items: ['Profile', 'Sign Out'],
    genericMenu: [
      { icon: 'mdi-view-dashboard', text: 'Personal Dashboard', route: '/' },
      {
        icon: 'mdi-silverware',
        text: 'Food',
        route: '/food',
        submenu: [
          {
            icon: 'mdi-calendar-today',
            text: 'Meal Diary',
            route: '/food/diary',
          },
          {
            icon: 'mdi-cup-water',
            text: 'Water Intake',
            route: '/food/fluids',
          },
          {
            icon: 'mdi-archive',
            text: 'Inventory',
            route: '/food/inventory',
          },
        ],
      },
      {
        icon: 'mdi-weight-kilogram',
        text: 'Weight Tracking',
        route: '/measures',
      },
      {
        icon: 'mdi-calendar',
        text: 'Calendar',
        route: '/calendar',
      },
      {
        icon: 'mdi-run-fast',
        text: 'Activities (soon!)',
        route: '/comingsoon',
      },
    ],
  }),

  computed: {
    LoggedIn() {
      return this.$store.state.auth.loggedIn
    },

    Avatar() {
      if (this.$store.state.auth.user.user_settings.avatar) {
        return this.$store.state.auth.user.user_settings.avatar
      }
    },

    NameAvatar() {
      return this.$store.state.auth.user.username.charAt(0)
    },
  },

  methods: {
    logOut() {
      this.$auth.logout()
      this.$router.push('/login')
    },
  },
}
</script>

<style></style>
