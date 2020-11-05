<template>
  <div>
    <section class="performance-facts">
      <header class="performance-facts__header">
        <h1 class="performance-facts__title">Nutrition Facts</h1>
        <p>Serving Size 1/2 cup (about 82g)</p>
        <p>Serving Per Container 8</p>
      </header>
      <table class="performance-facts__table">
        <thead>
          <tr>
            <th colspan="3" class="small-info">
              Amount Per Serving
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th colspan="2">
              <b>Calories</b>
              {{ product.product.nutriments.energy_value }}
            </th>
          </tr>
          <tr class="thick-row">
            <td colspan="3" class="small-info">
              <b>100ml / 100g</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Total Fat</b>
            </th>
            <td>
              <b>
              {{ product.product.nutriments.fat }}g
              </b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Saturated Fat {{ product.product.nutriments["saturated-fat"] }}g
            </th>
            <td>
              <b>22%</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Cholesterol</b>
              55mg
            </th>
            <td>
              <b>18%</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Sodium</b>
              40mg
            </th>
            <td>
              <b>2%</b>
            </td>
          </tr>
          <tr>
            <th colspan="2">
              <b>Total Carbohydrate</b>
              17g
            </th>
            <td>
              <b>6%</b>
            </td>
          </tr>
          <tr>
            <td class="blank-cell"></td>
            <th>
              Dietary Fiber 1g
            </th>
            <td>
              <b>4%</b>
            </td>
          </tr>
          <tr v-if="product.product.nutriments.sugars">
            <td class="blank-cell"></td>
            <th>
              Sugars 14g
            </th>
            <td></td>
          </tr>
          <tr v-if="product.product.nutriments.proteins" class="thick-end">
            <th colspan="2">
              <b>Protein</b>
            </th>
            <td>{{ product.product.nutriments.proteins }}g</td>
          </tr>
        </tbody>
      </table>
    </section>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  name: 'NutritionFacts',

  props: {
    product: {
      type: Object
    }
  },

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
