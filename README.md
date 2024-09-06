# Uy vazifasi: Fitness kuzatuv dasturi uchun Foydalanuvchilar API-ni to'liq amalga oshirish

## Maqsad
Bu vazifaning maqsadi fitness kuzatuv dasturining `Users API` qismini parolni tiklash funksionalligi bilan to‘liq amalga oshirishdir. Sizlar parolni xavfsiz tarzda tiklash, `CRUD` operatsiyalarini bajarish va `JWT` asosida autentifikatsiya qilishlari kerak.

## Vazifa Ko'rsatmalari:
1. **CRUD handlerlarni to'liq qilish - foydalanuvchi yozuvlarini yaratish, olish, yangilash va o‘chirish uchun.**

2. **Parol hashingini amalga oshirish - ilgari yaratilgan yordamchi funksiyalar yordamida parolni hashing qilish.**

3. **Parolni tiklash mexanizmini amalga oshirish**
    1. **Parolni tiklash so'rovini yaratish**
        - Foydalanuvchi email manzilini kiritadi.
        - Server email orqali yuborish uchun vaqtinchalik `token` yaratadi.
        - Ushbu `token` foydalanuvchi `reset password` so'rovi bilan bog‘liq qilib ma'lumotlar bazasida saqlanadi.
        - Foydalanuvchiga email orqali parolni tiklash havolasi yuboriladi.
    2. **Tokenni tekshirish**:
        - Foydalanuvchi emailda yuborilgan `token` yordamida parolini tiklash so'rovini amalga oshiradi.
        - Server tokenni tekshiradi va agar yaroqli bo‘lsa, foydalanuvchi yangi parolni kiritadi.
    3. **Yangi parolni saqlash:**:
        - Foydalanuvchi kiritgan yangi parol `hashing` qilinadi va foydalanuvchi yozuvi yangilanadi.

4. **JWT asosida autentifikatsiya qo‘shish - foydalanuvchilarning sezgir routelar ni himoya qilish.**:

5. **Foydalanuvchi handlerlar uchun unit testlarni yozish va barcha funksionalliklarni to‘g‘ri ishlashini ta’minlash.**




