package main

// Чем отличаются RWMutex от Mutex?

// answer:
// RWMutex имеет дополнительно два метода, RLock & RUnlock.
// что позволяет читать данные переменной из другой горутины
