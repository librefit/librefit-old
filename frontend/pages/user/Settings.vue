<template>
  <v-container fluid>

    <base-material-card>
      <template v-slot:heading>User Profile</template>
      <v-card-text>
        <p>Full Name: </p>
        <p>Email: </p>
        <p>Date of Birth: </p>
      </v-card-text>
    </base-material-card>
    
    <div class="py-7" />
    
    <base-material-card>
      <template v-slot:heading>Date and Time</template>
      <v-card-text>
        <p>Date Format</p>
        <p>Calendar Week Starts On: </p>
        <p>Time Format: </p>
      </v-card-text>
    </base-material-card>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    cards: ['Today', 'Yesterday'],
    drawer: null,
    links: [
      ['mdi-inbox-arrow-down', 'Inbox'],
      ['mdi-send', 'Send'],
      ['mdi-delete', 'Trash'],
      ['mdi-alert-octagon', 'Spam']
    ]
  })
}
</script>
