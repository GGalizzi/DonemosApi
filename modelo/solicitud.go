package modelo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Solicitud struct {
	SolicitudId     bson.ObjectId   `json:"solicitudID" bson:"_id"`
	UsuarioId       int             `json:"usuarioID" bson:"usuarioID"`
	Fecha           time.Time       `json:"fechaCreacion" bson:"fecha"`
	Vigente         bool            `json:"estaVigente" bson:"estaVigente"`
	CantidadDadores int             `json:"cantidadDadores" bson:"cantidadDadores"`
	NombrePaciente  string          `json:"nombrePaciente" bson:"nombrePaciente"`
	Institucion     string          `json:"institucion" bson:"institucion"`
	Direccion       string          `json:"direccion" bson:"direccion"`
	HoraDesde       string          `json:"horaDesde" bson:"horaDesde"`
	HoraHasta       string          `json:"horaHasta" bson:"horaHasta"`
	Adicionales     string          `json:"datosAdicionales" bson:"datosAdicionales"`
	Grupo           GrupoSanguineo  `json:"grupoSanguineo" bson:"grupoSanguineo"`
	Factor          FactorSanguineo `json:"factorSanguineo" bson:"factorSanguineo"`
	provinciaId     bson.ObjectId
	ciudadId        bson.ObjectId
}
type Solicitudes []Solicitud

func NewSolicitud(nombre, institucion string) *Solicitud {
	s := new(Solicitud)
	s.NombrePaciente = nombre
	s.Institucion = institucion
	return s
}

func (s *Solicitud) GetId() bson.ObjectId {
	return s.SolicitudId
}

func (s *Solicitud) SetIdHex(hex string) {
	s.SolicitudId = bson.ObjectIdHex(hex)
}

func (s *Solicitud) SetId(id bson.ObjectId) {
	s.SolicitudId = id
	// TODO: TEMP: These Ids will be set from the received body
	s.provinciaId = bson.NewObjectId()
	s.ciudadId = bson.NewObjectId()

}

// func (s *Solicitud) MarshalJSON() ([]byte, error) {
// 	type Alias Solicitud
// 	return json.Marshal(&struct{
// 		Provincia bson.ObjectId `json:""`
// 	})
// }
