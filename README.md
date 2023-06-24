Trong kiến trúc Clean Architecture, chúng ta phân ra làm 3 tầng
+ Biz (Business Layer)
+ Storage Layer
+ Transport Layer

Tầng Business và tầng Storage không hiểu nhau và không có liên hệ.
Chúng ta cần sử dụng tầng Transport như phương tiện giao tiếp giữa chúng.

## Thông qua các interface.