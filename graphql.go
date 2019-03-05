package util

// GQLBuild : ignore identical field
func (s Str) GQLBuild(typename, field, fieldtype string) (gql string) {
	if ok, pos := s.LooseSearchStrs("type", typename, "{", " \t"); ok {
		content, _, r := s[pos:].BracketsPos(BCurly, 1, 1)
		if sCtn(content, field+":") {
			return s.V()
		}
		gql = s.V()[:pos+r]
		tail := s.V()[pos+r+1:]
		add := fSf("\t%s: %s\n}", field, fieldtype)
		gql += add + tail
	} else {
		s += Str(fSf("\n\ntype %s {\n\t%s: %s\n}", typename, field, fieldtype))
		gql = s.V()
	}
	return gql
}
