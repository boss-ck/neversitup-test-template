# Project

## Project Structure

├── src
│ ├── domain
│ │ └── entity.go
│ ├── usecases
│ │ └── usecase.go
│ ├── interface
│ │ ├── controller
│ │ │ └── controller.go
│ │ ├── repository
│ │ │ └── repository.go
│ ├── router
│ │ └── router.go
│ └── config
│ └── config.go
├── pkg # แพ็คเกจต่างๆที่ใช้
|── main.go
├── go.mod
└── go.sum

### Conventions

1. ใช้ชื่อไฟล์และโฟลเดอร์ให้สื่อความหมายและเป็นมาตรฐาน เช่น `controllers`, `repositories`, `usecases`
2. ใช้ชื่อตัวแปรและฟังก์ชันให้สื่อความหมายและเป็นมาตรฐาน เช่น ใช้ `camelCase` สำหรับตัวแปรและฟังก์ชันใน Go ที่ไม่ได้มี export
3. ใช้คอมเมนต์อธิบายโค้ดเพื่อเพิ่มความเข้าใจและความชัดเจน
4. ใช้การจัดกลุ่มโค้ดที่เกี่ยวข้องเข้าด้วยกันในโฟลเดอร์หรือแพ็คเกจที่เหมาะสม
5. ใช้การเขียน unit tests เพื่อทดสอบโค้ดและรักษาความถูกต้องของโปรแกรม
