<template>
  <div>
    <button class="btn btn-primary" @click="fetchRecords">Obtener registros</button>
    <table v-if="records.length" class="table table-dark mt-3">
      <thead>
        <tr>
          <th>Name</th>
          <th>Album</th>
          <th>Year</th>
          <th>Rank</th>
          <th>Date</th>
          <th>Hour</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="record in records" :key="record._id">
          <td>{{ record.name }}</td>
          <td>{{ record.album }}</td>
          <td>{{ record.year }}</td>
          <td>{{ record.rank }}</td>
          <td>{{ record.fecha }}</td>
          <td>{{ record.hora }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else class="mt-3">
      No hay registros para mostrar
    </div>
  </div>
</template>

<script>
import axios from 'axios'; // Importa la biblioteca Axios

export default {
  data() {
    return {
      records: [] // Almacena los registros recuperados de la base de datos
    };
  },
  methods: {
    fetchRecords() {
      axios.get('/votos') // Realiza una solicitud GET al endpoint de tu backend
        .then(response => {
          this.records = response.data; // Actualiza la lista de registros con los datos recibidos del backend
        })
        .catch(error => {
          console.error('Error al obtener registros:', error);
        });
    }
  }
};
</script>

<style>
/* No es necesario agregar estilos de Bootstrap aquí, ya que las clases de Bootstrap se aplican directamente en las plantillas */
</style>
