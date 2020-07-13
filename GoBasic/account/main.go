// 동기화 문제
// lock을 걸어서 해결, mutex를 이용 
// lock을 사용하는 것은 매우 복잡한 문제, 어느 범위로 lock을 걸 것인가? 
// 멀티태스킹을 한다는 것은 그래서 복잡한 문제임
// golang은 channel을 통해 이를 쉽게 접근할 수 있게 도와줌. but 정답은 아님
// 상황에 따라 lock을 쓰거나 channel을 사용해야함 

package main

import (
 "fmt"
 "math/rand"
 "sync"
 "time"
)

type Account struct {
 balance int
 mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
 a.mutex.Lock()
 a.balance -= val
 a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
 a.mutex.Lock()
 a.balance += val
 a.mutex.Unlock()
}

func (a *Account) Balance() int {
 a.mutex.Lock()
 balance := a.balance
 a.mutex.Unlock()
 return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
 globalLock.Lock()
 accounts[sender].Widthdraw(money)
 accounts[receiver].Deposit(money)
 globalLock.Unlock()
}

func GetTotalBalance() int {
 globalLock.Lock()
 total := 0
 for i := 0; i < len(accounts); i++ {
  total += accounts[i].Balance()
 }
 globalLock.Unlock()
 return total
}

func RandomTranfer() {
 var sender, balance int
 for {
  sender = rand.Intn(len(accounts))
  balance = accounts[sender].Balance()
  if balance > 0 {
   break
  }
 }

 var receiver int
 for {
  receiver = rand.Intn(len(accounts))
  if sender != receiver {
   break
  }
 }

 money := rand.Intn(balance)
 Transfer(sender, receiver, money)
}

func GoTransfer() {
 for {
  RandomTranfer()
 }
}

func PrintTotalBalance() {
 fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
 for i := 0; i < 20; i++ {
  accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
 }
 globalLock = &sync.Mutex{}

 PrintTotalBalance()

 for i := 0; i < 10; i++ {
  go GoTransfer()
 }

 for {
  PrintTotalBalance()
  time.Sleep(100 * time.Millisecond)
 }
}