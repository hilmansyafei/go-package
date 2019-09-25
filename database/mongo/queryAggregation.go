package mongo

import "github.com/globalsign/mgo/bson"

// QueryAggregate : struct for query aggregation.
type QueryAggregate struct {
	Query []bson.M
}

// AddWhere : add condition to query.
func (q *QueryAggregate) AddWhere(where bson.M) error {
	matchQuery := map[string]interface{}{
		"$match": where,
	}
	q.Query = append(q.Query, matchQuery)
	return nil
}

// JoinCollection : join to another collection.
func (q *QueryAggregate) JoinCollection(joinCollection string, joinField string, mainField string, as string) error {
	lookupQuery := map[string]interface{}{
		"$lookup": bson.M{
			"from":         joinCollection,
			"foreignField": joinField,
			"localField":   mainField,
			"as":           as,
		},
	}
	q.Query = append(q.Query, lookupQuery)
	unwind := map[string]interface{}{
		"$unwind": bson.M{
			"path":                       "$" + joinCollection,
			"preserveNullAndEmptyArrays": true,
		},
	}
	q.Query = append(q.Query, unwind)
	return nil
}

// Field : list field to show.
func (q *QueryAggregate) Field(fields bson.M) error {
	projectQuery := map[string]interface{}{
		"$project": fields,
	}
	q.Query = append(q.Query, projectQuery)
	return nil
}

// Sort : add sort condition to query.
func (q *QueryAggregate) Sort(sorts bson.M) error {
	sortQuery := map[string]interface{}{
		"$sort": sorts,
	}
	q.Query = append(q.Query, sortQuery)
	return nil
}

// Pagination : add offset and limit to query.
func (q *QueryAggregate) Pagination(offset int, limit int) error {
	skip := map[string]interface{}{
		"$skip": offset,
	}
	q.Query = append(q.Query, skip)
	limitQuery := map[string]interface{}{
		"$limit": limit,
	}
	q.Query = append(q.Query, limitQuery)
	return nil
}
