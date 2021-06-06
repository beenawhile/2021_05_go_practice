package main

// O.O.D : Object Oriented Design
// - object 중심의 설계 방법

// 설계가 왜 중요한가?
//  -> 혼자 프로그래밍 하는 것이 아니기 때문

// S.O.L.I.D : object 중심의 설계에서 지향해야할 5가지 목표
// - 이상향에 가까움
// - 달성해야하는 목표가 아님. 지향해야할 목표.
// - 실무에서는 항상 현실적인 문제가 있어 trade-off가 있어야 하지만, 아예 무시해서는 안됨.

//  1. Single Responsibility Principle : 단일 책임 원칙
//    - 가장 단순하면서 가장 중요한 원칙
//    - 단순하지만 지키기 어렵고, 다 지켰다고해서 좋은 코드도 아님
//    - 하지만 이렇게 가는게 좋다. 노력할 필요가 있다
//    - 하나의 객체는 하나의 책임을 가져야 한다
//    - ex. 예금 잔고 객체는 입금, 출금 (혹은 하나의 책임으로 봤을 때 입출금)의 책임을 가지고 있을 때 이를 하나의 책임으로 봐야하는가 여러개의 책임으로 봐야하는가

//  2. Open Closed Principle
//    - 확장에는 열려있고 변경에는 닫혀있음
//    - 다섯가지 규칙이 정의된 이유라고 봐도 과언이 아님 => 최종 목표 : 의존성 낮추고 응집성 높임

//  3. Liskov Subsititution Principle : 리스코프 치환 이론
//    - OOD 원칙 중 가장 어렵다 얘기함
//    - O(x) x는 T의 instance, O(y) y는 T를 확장한 S의 instance 일 때, O(x)와 O(y)는 같은 방식으로 동작해야함
//     => base type의 기본 동작을 확장한 type에서 바꾸지 마라는 의미
//    - go lang에서는 상속을 지원하지 않기 때문에, base type을 바꿀 일이 없기 때문에 크게 신경 쓰지 않아도 됨

//  4. Interface Segregation : 인터페이스 분리 원칙
//    - 여러개의 관계를 모아놓은 interface 보다 관계 하나씩 정의하는 것이 더 좋다
//    - 클라이언트가 자신이 이용하지 않는 메서드에 의존하지 않아야 한다는 원칙

//  5. Dependency Inversion Principle : 의존성 역전 법칙
//    - 관계는 인터페이스에 의존해야지 객체에 의존하면 안된다

// 1. Single Responsibility Principle Example
// 2. Open Closed Principle 에도 적용됨

// 잘못된 설계
// -----------------------------------
// type FinanceReport struct {}

// func (r *FinanceReport) MakeReport() *Report {}
// func (r *FinanceReport) SendReport(email string) *Report {}

// 만약 SendReport 에 email이 아니라 http, network 등으로 바뀌면? 계속 추가해줘야함 => 설계가 잘못됨
// -----------------------------------

// 재설계
// -----------------------------------
// type FinanceReport struct {}

// func (r *FinanceReport) MakeReport () *Report {}

// type ReportSender interface {
// 	SendReport(*Report)
// }

// type EmailReportSender struct {}

// func (s *EmailReportSender) SendReport(r *Report) {}

// type FileReportSender struct {}

// func (s *FileReportSender) SendReport(r *Report) {}
// -----------------------------------

// 4. Interface Segregation Principle

// type Actor interface {
// 	Move()
// 	Attack()
// 	Talk()
// }

// func MoveTo(a Actor) {
// 	a.Move()
// 	a.Attack()
// }

// -----------------------------------
// - 다음과 같이 한 함수에서 2개의 기능을 동시에 하는 것은 Single Responsibility 원칙에 위배됨
//   => interface를 분할하면 이러한 것을 미연에 방지할 수 있게 됨
// - 이런식으로 하나의 interface에서 다양한 관계를 가지고 있는거보다 따로 분리하는 것이 낫다

// -----------------------------------
// type Talkable interface {
// 	Talk()
// }
// type Attackable interface {
// 	Attack()
// }
// type Movable interface {
// 	Move()
// }

// func MoveToWhere(a Movable) {
// 	a.Move()
// }

// -----------------------------------
// - 의존성을 떨어뜨리게됨

// 5. Dependency Inversion Principle
// -----------------------------------
// type Player struct{}

// type Monster struct{}

// func (p *Player) Attack(m *Monster) {}

// 플레이어가 플레이어는 못때리는가?
// func (p *Player) Attack(m *Player) {}

// 몬스터가 플레이어는 못때리는가?
// func (p *Monster) Attack(m *Player) {}

// 몬스터가 몬스터는 못때리는가?
// func (p *Monster) Attack(m *Monster) {}

//  이런 식으로 객체가 늘어나면 새로 추가 해줘야함 => OCP(Open Closed Principle)에 어긋남
//     확장에는 열려있고 변경에는 닫혀있어야하는데 변경에 닫혀있지 못함
// -----------------------------------

// -----------------------------------
// type Attackable interface {
// 	Attack(BeAttackable)
// }
// type BeAttackable interface {
// 	BeAttacked()
// }

// func Attack(attacker *Attackable, defender *BeAttackable) {
// 	attacker.Attack(defender)
// }

// func (p *Player) Attack(target *BeAttackable) {}

// func (m *Monster) Attack(target * BeAttackable){}
// -----------------------------------
