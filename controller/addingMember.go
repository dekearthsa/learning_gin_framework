package controller

import "learning_gin_framework/model"

func AddingMember() []model.Member {
	var setMember []model.Member
	setAge := 0
	for i := 0; i < 10; i++ {
		setAge = setAge + 1
		setMember = append(setMember, model.Member{Name: "EE", Age: setAge, Phone: "32233244"})
	}

	return setMember
}
