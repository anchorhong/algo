package _721

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	var accountIDs []int
	var accountWeights []int
	var accountEmails []emailSet
	for i, account := range accounts {
		accountIDs = append(accountIDs, i)
		accountWeights = append(accountWeights, 1)
		es := make(emailSet)
		es.addAll(account[1:])
		accountEmails = append(accountEmails, es)
	}
	uf := unionFind{
		id:     accountIDs,
		emails: accountEmails,
		wt:     accountWeights,
		count:  len(accounts),
	}
	emailMapping := make(map[string]int)
	for curID, account := range accounts {
		emails := account[1:]
		for _, email := range emails {
			if otherID, exists := emailMapping[email]; exists {
				uf.union(otherID, curID)
			} else {
				emailMapping[email] = curID
			}
		}
	}
	var res [][]string
	for i, _ := range uf.id {
		root := uf.find(i)
		if root == i {
			tmpRes := []string{accounts[i][0]}
			es := uf.emails[i].toSlice()
			sort.Strings(es)
			tmpRes = append(tmpRes, es...)
			res = append(res, tmpRes)
		}
	}
	return res
}

type emailSet map[string]interface{}

func (s *emailSet) add(email string) {
	(*s)[email] = nil
}

func (s *emailSet) addAll(emails []string) {
	for _, email := range emails {
		(*s)[email] = nil
	}
}

func (s *emailSet) merge(set emailSet) {
	for k, _ := range set {
		s.add(k)
	}
}

func (s *emailSet) toSlice() []string {
	var emails []string
	for k, _ := range *s {
		emails = append(emails, k)
	}
	return emails
}

func (s *emailSet) contains(email string) bool {
	if _, ok := (*s)[email]; ok {
		return true
	}
	return false
}

type unionFind struct {
	id     []int
	emails []emailSet
	wt     []int
	count  int
}

func (u unionFind) find(p int) int {
	if u.id[p] != p {
		u.id[p] = u.find(u.id[p])
	}
	return u.id[p]
}

func (u unionFind) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.wt[pRoot] < u.wt[qRoot] {
		u.id[pRoot] = qRoot
		u.wt[qRoot] += u.wt[pRoot]
		u.emails[qRoot].merge(u.emails[pRoot])
	} else {
		u.id[qRoot] = pRoot
		u.wt[pRoot] += u.wt[qRoot]
		u.emails[pRoot].merge(u.emails[qRoot])
	}
	u.count--
}
