<template>
  <div>
    <section class="performance-facts">
      <header class="performance-facts__header">
        <h1 class="performance-facts__title">Nutrition Facts</h1>
      </header>
      <table class="performance-facts__table">
        <tbody>
          <tr>
            <th colspan="2" class="text-h5">
              <b>Calories</b>
            </th>
            <td class="text-h5">
              <b>{{ product.calories }} </b>
            </td>
          </tr>
          <tr class="thick-row">
            <td colspan="3" class="small-info">
              <b>100 {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Total Fat</b>
            </th>
            <td>
              <b> {{ product.fat_total }} {{ unit }} </b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Saturated Fat
            </th>
            <td>
              <b>{{ product.fat_saturated }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Polyunsaturated Fat
            </th>
            <td>
              <b>{{ product.fat_polyunsaturated }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Monounsaturated Fat
            </th>
            <td>
              <b>{{ product.fat_monounsaturated }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Trans
            </th>
            <td>
              <b>{{ product.fat_trans }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Cholesterol</b>
            </th>
            <td>
              <b>{{ product.fat_cholesterol }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Sodium</b>
            </th>
            <td>
              <b>{{ product.sodium }} {{ unit }}</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Potassium</b>
            </th>
            <td>
              <b>{{ product.potassium }} {{ unit }}</b>
            </td>
          </tr>

          <tr>
            <th colspan="2">
              <b>Total Carbohydrate</b>
            </th>
            <td>
              <b>{{ product.carbs }}</b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Sugars
            </th>
            <td>
              <b>{{ product.sugars }} {{ unit }}</b>
            </td>
          </tr>
          <tr class="thick-end">
            <th colspan="2">
              <b>Protein</b>
            </th>
            <td>{{ product.proteins }} {{ unit }}</td>
          </tr>
        </tbody>
      </table>
      <p class="text-center">
        <a
          :href="'https://world.openfoodfacts.org/product/' + product.off_code"
          target="_blank"
          >More details in OpenFoodFacts</a
        >
      </p>
    </section>
  </div>
</template>

<script>
export default {
  name: 'NutritionFactsInventory',

  data: () => ({
    unit: null
  }),

  props: {
    product: {
      type: Object
    }
  },

  computed: {
    unitType() {
      this.unit = 'g'
    }
  }
}
</script>

<style lang="scss">
body {
  font-size: small;
  line-height: 1.4;
}
p {
  margin: 0;
}

.performance-facts {
  border: 1px solid black;
  width: 280px;
  padding: 0.5rem;
  table {
    border-collapse: collapse;
  }
}
.performance-facts__title {
  font-weight: bold;
  font-size: 2rem;
  margin: 0 0 0.25rem 0;
}
.performance-facts__header {
  border-bottom: 10px solid black;
  padding: 0 0 0.25rem 0;
  margin: 0 0 0.5rem 0;
  p {
    margin: 0;
  }
}
.performance-facts__table {
  width: 100%;
  thead tr {
    th,
    td {
      border: 0;
    }
  }
  th,
  td {
    font-weight: normal;
    text-align: left;
    padding: 0.25rem 0;
    border-top: 1px solid black;
    white-space: nowrap;
  }
  td {
    &:last-child {
      text-align: right;
    }
  }
  .blank-cell {
    width: 1rem;
    border-top: 0;
  }
  .thick-row {
    th,
    td {
      border-top-width: 5px;
    }
  }
}

.small-info {
  font-size: 0.7rem;
}

.performance-facts__table--small {
  @extend .performance-facts__table;
  border-bottom: 1px solid #999;
  margin: 0 0 0.5rem 0;
  thead {
    tr {
      border-bottom: 1px solid black;
    }
  }
  td {
    &:last-child {
      text-align: left;
    }
  }
  th,
  td {
    border: 0;
    padding: 0;
  }
}

.performance-facts__table--grid {
  @extend .performance-facts__table;
  margin: 0 0 0.5rem 0;
  td {
    &:last-child {
      text-align: left;
      &::before {
        content: '•';
        font-weight: bold;
        margin: 0 0.25rem 0 0;
      }
    }
  }
}

.text-center {
  text-align: center;
}
.thick-end {
  border-bottom: 10px solid black;
}
.thin-end {
  border-bottom: 1px solid black;
}
</style>
