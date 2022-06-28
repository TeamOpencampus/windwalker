// Code generated by entc, DO NOT EDIT.

package company

import (
	"windwalker/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CompanyName applies equality check predicate on the "company_name" field. It's identical to CompanyNameEQ.
func CompanyName(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompanyName), v))
	})
}

// ContactPersonName applies equality check predicate on the "contact_person_name" field. It's identical to ContactPersonNameEQ.
func ContactPersonName(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonEmail applies equality check predicate on the "contact_person_email" field. It's identical to ContactPersonEmailEQ.
func ContactPersonEmail(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonPhone applies equality check predicate on the "contact_person_phone" field. It's identical to ContactPersonPhoneEQ.
func ContactPersonPhone(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonPhone), v))
	})
}

// CompanyNameEQ applies the EQ predicate on the "company_name" field.
func CompanyNameEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompanyName), v))
	})
}

// CompanyNameNEQ applies the NEQ predicate on the "company_name" field.
func CompanyNameNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompanyName), v))
	})
}

// CompanyNameIn applies the In predicate on the "company_name" field.
func CompanyNameIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCompanyName), v...))
	})
}

// CompanyNameNotIn applies the NotIn predicate on the "company_name" field.
func CompanyNameNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCompanyName), v...))
	})
}

// CompanyNameGT applies the GT predicate on the "company_name" field.
func CompanyNameGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompanyName), v))
	})
}

// CompanyNameGTE applies the GTE predicate on the "company_name" field.
func CompanyNameGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompanyName), v))
	})
}

// CompanyNameLT applies the LT predicate on the "company_name" field.
func CompanyNameLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompanyName), v))
	})
}

// CompanyNameLTE applies the LTE predicate on the "company_name" field.
func CompanyNameLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompanyName), v))
	})
}

// CompanyNameContains applies the Contains predicate on the "company_name" field.
func CompanyNameContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCompanyName), v))
	})
}

// CompanyNameHasPrefix applies the HasPrefix predicate on the "company_name" field.
func CompanyNameHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCompanyName), v))
	})
}

// CompanyNameHasSuffix applies the HasSuffix predicate on the "company_name" field.
func CompanyNameHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCompanyName), v))
	})
}

// CompanyNameEqualFold applies the EqualFold predicate on the "company_name" field.
func CompanyNameEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCompanyName), v))
	})
}

// CompanyNameContainsFold applies the ContainsFold predicate on the "company_name" field.
func CompanyNameContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCompanyName), v))
	})
}

// ContactPersonNameEQ applies the EQ predicate on the "contact_person_name" field.
func ContactPersonNameEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameNEQ applies the NEQ predicate on the "contact_person_name" field.
func ContactPersonNameNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameIn applies the In predicate on the "contact_person_name" field.
func ContactPersonNameIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContactPersonName), v...))
	})
}

// ContactPersonNameNotIn applies the NotIn predicate on the "contact_person_name" field.
func ContactPersonNameNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContactPersonName), v...))
	})
}

// ContactPersonNameGT applies the GT predicate on the "contact_person_name" field.
func ContactPersonNameGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameGTE applies the GTE predicate on the "contact_person_name" field.
func ContactPersonNameGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameLT applies the LT predicate on the "contact_person_name" field.
func ContactPersonNameLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameLTE applies the LTE predicate on the "contact_person_name" field.
func ContactPersonNameLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameContains applies the Contains predicate on the "contact_person_name" field.
func ContactPersonNameContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameHasPrefix applies the HasPrefix predicate on the "contact_person_name" field.
func ContactPersonNameHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameHasSuffix applies the HasSuffix predicate on the "contact_person_name" field.
func ContactPersonNameHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameEqualFold applies the EqualFold predicate on the "contact_person_name" field.
func ContactPersonNameEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonNameContainsFold applies the ContainsFold predicate on the "contact_person_name" field.
func ContactPersonNameContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContactPersonName), v))
	})
}

// ContactPersonEmailEQ applies the EQ predicate on the "contact_person_email" field.
func ContactPersonEmailEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailNEQ applies the NEQ predicate on the "contact_person_email" field.
func ContactPersonEmailNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailIn applies the In predicate on the "contact_person_email" field.
func ContactPersonEmailIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContactPersonEmail), v...))
	})
}

// ContactPersonEmailNotIn applies the NotIn predicate on the "contact_person_email" field.
func ContactPersonEmailNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContactPersonEmail), v...))
	})
}

// ContactPersonEmailGT applies the GT predicate on the "contact_person_email" field.
func ContactPersonEmailGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailGTE applies the GTE predicate on the "contact_person_email" field.
func ContactPersonEmailGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailLT applies the LT predicate on the "contact_person_email" field.
func ContactPersonEmailLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailLTE applies the LTE predicate on the "contact_person_email" field.
func ContactPersonEmailLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailContains applies the Contains predicate on the "contact_person_email" field.
func ContactPersonEmailContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailHasPrefix applies the HasPrefix predicate on the "contact_person_email" field.
func ContactPersonEmailHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailHasSuffix applies the HasSuffix predicate on the "contact_person_email" field.
func ContactPersonEmailHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailEqualFold applies the EqualFold predicate on the "contact_person_email" field.
func ContactPersonEmailEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonEmailContainsFold applies the ContainsFold predicate on the "contact_person_email" field.
func ContactPersonEmailContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContactPersonEmail), v))
	})
}

// ContactPersonPhoneEQ applies the EQ predicate on the "contact_person_phone" field.
func ContactPersonPhoneEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneNEQ applies the NEQ predicate on the "contact_person_phone" field.
func ContactPersonPhoneNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneIn applies the In predicate on the "contact_person_phone" field.
func ContactPersonPhoneIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContactPersonPhone), v...))
	})
}

// ContactPersonPhoneNotIn applies the NotIn predicate on the "contact_person_phone" field.
func ContactPersonPhoneNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContactPersonPhone), v...))
	})
}

// ContactPersonPhoneGT applies the GT predicate on the "contact_person_phone" field.
func ContactPersonPhoneGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneGTE applies the GTE predicate on the "contact_person_phone" field.
func ContactPersonPhoneGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneLT applies the LT predicate on the "contact_person_phone" field.
func ContactPersonPhoneLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneLTE applies the LTE predicate on the "contact_person_phone" field.
func ContactPersonPhoneLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneContains applies the Contains predicate on the "contact_person_phone" field.
func ContactPersonPhoneContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneHasPrefix applies the HasPrefix predicate on the "contact_person_phone" field.
func ContactPersonPhoneHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneHasSuffix applies the HasSuffix predicate on the "contact_person_phone" field.
func ContactPersonPhoneHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneEqualFold applies the EqualFold predicate on the "contact_person_phone" field.
func ContactPersonPhoneEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContactPersonPhone), v))
	})
}

// ContactPersonPhoneContainsFold applies the ContainsFold predicate on the "contact_person_phone" field.
func ContactPersonPhoneContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContactPersonPhone), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.JobPost) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		p(s.Not())
	})
}