package mock

import (
	"strings"
	"testing"

	"github.com/milosgajdos/kraph/api"
	"github.com/milosgajdos/kraph/api/generic"
	"github.com/milosgajdos/kraph/query"
)

func newTestTop() *Top {
	top := NewTop()

	for resName, meta := range Resources {
		groups := ResourceData[resName]["groups"]
		versions := ResourceData[resName]["versions"]
		for _, group := range groups {
			for _, version := range versions {
				gv := strings.Join([]string{group, version}, "/")
				if gvObject, ok := ObjectData[gv]; ok {
					kind := meta["kind"]
					ns := meta["ns"]
					if len(ns) == 0 {
						ns = api.NsNan
					}

					nsKind := strings.Join([]string{ns, kind}, "/")

					if names, ok := gvObject[nsKind]; ok {
						for _, name := range names {
							uid := strings.Join([]string{ns, kind, name}, "/")
							links := make(map[string]api.Relation)
							if rels, ok := ObjectLinks[uid]; ok {
								for obj, rel := range rels {
									links[obj] = generic.NewRelation(rel)
								}
							}
							object := generic.NewObject(name, kind, ns, generic.NewUID(uid), links)
							top.Add(object)
						}
					}
				}
			}
		}
	}

	return top
}

func TestObjects(t *testing.T) {
	top := newTestTop()

	objects := top.Objects()
	if len(objects) == 0 {
		t.Errorf("no objects found")
	}
}

func TestGetUID(t *testing.T) {
	top := newTestTop()

	for _, nsKinds := range ObjectData {
		for nsKind, names := range nsKinds {
			nsplit := strings.Split(nsKind, "/")
			ns, kind := nsplit[0], nsplit[1]
			for _, name := range names {
				uid := strings.Join([]string{ns, kind, name}, "/")
				objects, err := top.Get(query.UID(uid))
				if err != nil {
					t.Errorf("error getting object: %s: %v", uid, err)
					continue
				}

				if len(objects) != 1 {
					t.Errorf("expected 1 object, got: %d", len(objects))
					continue
				}

				if objects[0].UID().String() != uid {
					t.Errorf("expected object %s, got: %s", uid, objects[0].UID())
				}
			}
		}
	}
}

func TestGetNsKind(t *testing.T) {
	top := newTestTop()

	for _, nsKinds := range ObjectData {
		for nsKind, _ := range nsKinds {
			nsplit := strings.Split(nsKind, "/")
			ns, kind := nsplit[0], nsplit[1]
			objects, err := top.Get(query.Namespace(ns), query.Kind(kind))
			if err != nil {
				t.Errorf("error getting objects: %s/%s: %v", ns, kind, err)
				continue
			}

			for _, object := range objects {
				if object.Namespace() != ns || object.Kind() != kind {
					t.Errorf("expected: %s/%s, got: %s/%s", ns, kind, object.Namespace(), object.Kind())
				}
			}
		}
	}
}
