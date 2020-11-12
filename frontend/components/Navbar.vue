<template>
  <nav>
    <v-app-bar app clipped-left color="green">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <span class="title ml-3 mr-5"
        >Libre&nbsp;<span class="font-weight-light">Fit</span></span
      >
      <v-spacer></v-spacer>
      <v-spacer></v-spacer>
      <v-menu rounded="b-xl" offset-y>
        <template v-slot:activator="{ attrs, on }">
          <v-btn icon to="/login" v-bind="attrs" v-on="on">
            <v-icon>mdi-account</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item v-for="item in items" :key="item" link>
            <v-list-item-title v-text="item"></v-list-item-title>
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
        <!-- Only the logout buttom is not in the loop -->
        <v-list-item @click="logOut()">
          <v-list-item-icon><v-icon>mdi-logout</v-icon></v-list-item-icon>
          <v-list-item-content
            ><v-list-item-title>Log Out</v-list-item-title></v-list-item-content
          >
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      items: ['Profile', 'Sign Out'],
      genericMenu: [
        { icon: 'mdi-view-dashboard', text: 'Home', route: '/' },
        {
          icon: 'mdi-silverware',
          text: 'Food',
          route: '/food',
          submenu: [
            {
              icon: 'mdi-calendar-today',
              text: 'Meal Diary',
              route: '/food/diary'
            },
            {
              icon: 'mdi-cup-water',
              text: 'Water Intake',
              route: '/food/fluids'
            },
            {
              icon: 'mdi-archive',
              text: 'Inventory',
              route: '/food/inventory'
            }
          ]
        },
        {
          icon: 'mdi-weight-kilogram',
          text: 'Weight Tracking',
          route: '/measures'
        },
        {
          icon: 'mdi-calendar',
          text: 'Calendar',
          route: '/calendar'
        },
        {
          icon: 'mdi-run-fast',
          text: 'Activities (soon!)',
          route: '/comingsoon'
        },
        { divider: true },
        { icon: 'mdi-cog-outline', text: 'Settings', route: '/user/settings' }
      ]
    }
  },
  methods: {
    logOut() {
      this.$auth.logout()
      this.$router.push('/login')
    }
  }
}
</script>

<style></style>
