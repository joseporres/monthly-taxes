<script setup>
import { ref } from 'vue'
import { monthlyTaxes, logOut as AuhtLogOut } from '../api/index'
import { useRouter } from 'vue-router';
import { toast } from 'vue3-toastify';

const dni = ref('')
const router = useRouter();
const isLoading = ref(false);
const fullPage = ref(true);
const data = ref(null);


function onCancel() {
  isLoading.value = false;
}

async function getMonthlyTaxes() {

  try {
    data.value = null
    isLoading.value = true;
    let response = await monthlyTaxes(dni.value)
    let parseData = JSON.parse(response?.data?.response)
    if (parseData?.CoError === "9999") {
      // toast.error(parseData?.error)
      let message = parseData.Msg;
      toast.success(message)
      let dataJson = parseData?.dataJson
      let firma = 'data:image/jpg;base64, ' + dataJson.firma.replace("\n", '')
      let foto = 'data:image/jpg;base64, ' + dataJson.foto.replace("\n", '')
      let hizquierda = 'data:image/jpg;base64, ' + dataJson.hizquierda.replace("\n", '')
      let hderecha = 'data:image/jpg;base64, ' + dataJson.hderecha.replace("\n", '')
      let info = {
        ...dataJson.listaAni[0],
        imageFirma: firma,
        imageFoto: foto,
        imagehizquierda: hizquierda,
        imagehderecha: hderecha
      }
      data.value = info

    } else {
      let msgError = parseData?.Msg || 'Ocurrio un error'
      toast.error(msgError)
    }
    isLoading.value = false

  } catch (error) {
    console.log(error)
    toast.error('Ocurrio un error')
  }
}

async function logOut() {
  try {
    let response = await AuhtLogOut()
    console.log(response.data)
    localStorage.removeItem('token')
    router.push('/login')
  } catch (error) {
    console.log(error)
  }
}

</script>

<template>
  <div class="logout">
    <button class="btn btn-red" @click="logOut">Logout</button>
  </div>

  <div class="container">
    <h1 class="center title-sucess">Buscar por dni</h1>

    <div class="form-group">
      <input type="text" v-model="dni" 
       min="8" max="8"
      />

      <div class="center" v-if="!isLoading">
        <button class="btn btn-sucess" @click="getMonthlyTaxes"> Buscar</button>
      </div>
    </div>
  </div>

  <div v-if="isLoading" class="center">
    Cargando...
  </div>

  <div v-if="data" class="center" style="padding-top: 30px;">
    <div class="reniec-info">
      <table style='width:100%'>
        <tr style='border: 2px #E2E2EE solid'>
          <td style='width: 75%;'>
            <table style='width: 100%;'>
              <tr>
                <td style="width:50%" class="small bold">
                  Cód Único de indetificación:
                </td>
                <td style="width:50%" class="bold">
                  {{ data.nuDni }} - {{ data.digitoVerificacion }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Primer Apellido:
                </td>
                <td style="width:50%">
                  {{ data.apePaterno }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Segundo Apellido:
                </td>
                <td style="width:50%">
                  {{ data.apeMaterno }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Pre Nombres:
                </td>
                <td style="width:50%">
                  {{ data.preNombres }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Sexo:
                </td>
                <td style="width:50%">
                  {{ data.sexo }}
                </td>
              </tr>
              <tr>
                <td style="width: 50%;">
                  <table style="width: 100%;">
                    <tr>
                      <td style="width:50%" class="bold">
                        Nacimiento:
                      </td>
                      <td style="width:50%">
                        <table style="width: 100%;">
                          <tr class="bold">
                            <td>Fecha</td>
                          </tr>
                          <tr class="bold">
                            <td>Departamento</td>
                          </tr>
                          <tr class="bold">
                            <td>Provincia</td>
                          </tr>
                          <tr class="bold">
                            <td>Distrito</td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>

                </td>

                <td style="width: 50%">
                  <table style="width: 100%;">
                    <tr>
                      <td>
                        {{ data.feNacimiento }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.departamento }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.provincia }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.distrito }}
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Grado de Instrucción:
                </td>
                <td style="width:50%">
                  {{ data.gradoInstruccion }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Estado Civil:
                </td>
                <td style="width:50%">
                  {{ data.estadoCivil }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Estatura:
                </td>
                <td style="width:50%">
                  {{ data.estatura }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Fecha de Inscripción:
                </td>
                <td style="width:50%">
                  {{ data.feInscripcion }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Nombre del padre:
                </td>
                <td style="width:50%">
                  {{ data.nomPadre }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Nombre de la madre:
                </td>
                <td style="width:50%">
                  {{ data.nomMadre }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Fecha de emisión:
                </td>
                <td style="width:50%">
                  {{ data.feEmision }}
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Restricción:
                </td>
                <td style="width:50%">
                  {{ data.deRestriccion }}
                </td>
              </tr>
              <tr>
                <td style="width: 50%;">
                  <table style="width: 100%;">
                    <tr>
                      <td style="width:50%" class="bold">
                        Domicilio:
                      </td>
                      <td style="width:50%">
                        <table style="width: 100%;">
                          <tr class="bold">
                            <td>Departamento</td>
                          </tr>
                          <tr class="bold">
                            <td>Provincia</td>
                          </tr>
                          <tr class="bold">
                            <td>Distrito</td>
                          </tr>
                          <tr class="bold">
                            <td>Dirección</td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                  </table>

                </td>

                <td style="width: 50%">
                  <table style="width: 100%;">
                    <tr>
                      <td>
                        {{ data.depaDireccion }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.provDireccion }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.distDireccion }}
                      </td>
                    </tr>
                    <tr>
                      <td>
                        {{ data.desDireccion }}
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>
              <tr>
                <td style="width:50%" class="bold">
                  Fecha de caducidad:
                </td>
                <td style="width:50%">
                  {{ data.feCaducidad }}
                </td>
              </tr>
            </table>
          </td>
          <td style='width: 30%;'>
            <img :src="data.imageFoto" alt="" width="190">
            <img :src="data.imageFirma" alt="" width="210">
            <table>
              <tr>
                <td style="width:50%">
                  <img :src="data.imagehderecha" alt="foto3" width="100">
                </td>
                <td style="width:50%">
                  <img :src="data.imagehizquierda" alt="foto3" width="100">
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
    </div>

  </div>
</template>

<style scoped>
.reniec-info {
  width: 100%;
  border: 2px black solid;
  padding: 10px;
}

/*
  merdia srreen mayor a 1200
*/
@media screen and (min-width: 1200px) {
  .reniec-info {
    width: 50%;
    border: 2px #E2E2EE solid;
    padding: 10px;
  }
}
  


/* @media screen and (max-width: 600px) {
  .reniec-info {
    width: 50%;
    border: 2px #E2E2EE solid;
    padding: 10px;
  }
} */

.logout {
  display: flex;
  justify-content: flex-end;
}

.center {
  display: flex;
  justify-content: center;
}
</style>
