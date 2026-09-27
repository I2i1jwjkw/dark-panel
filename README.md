# Dark Panel

A dark, responsive, right-to-left starter UI for a VPN management panel.

یک رابط کاربری اولیه، واکنش‌گرا و راست‌چین برای پنل مدیریت VPN با تم تیره.

---

## English

### Installation / Run locally

**Option 1: Download**
1. Open this repository on GitHub.
2. Click **Code → Download ZIP**.
3. Extract the ZIP file on your computer or phone.
4. Open `index.html` in a modern web browser.

**Option 2: Clone with Git**
```bash
git clone https://github.com/I2i1jwjkw/dark-panel.git
cd dark-panel
```
Then open `index.html` in your browser. No build step or package installation is required for this prototype.

### Publish with GitHub Pages
1. Open the repository's **Settings**.
2. Select **Pages** in the sidebar.
3. Under **Build and deployment**, choose **Deploy from a branch**.
4. Select the `main` branch and `/(root)` folder, then click **Save**.
5. After GitHub Pages finishes deploying, use the URL shown in the Pages settings.

### Features
- Sample dashboard and statistics
- User list and search
- Add a demo user in the current browser page
- Servers, traffic, and settings screens
- Responsive layout for mobile and desktop

### Current status and limitations
This is a **frontend prototype only**. User records and statistics are sample data. Adding a user changes only the current page's in-memory data; it is not saved to a database. The panel is not connected to a VPN server, Xray, backend, or API and is **not ready for production use**.

### Planned next steps
1. Choose a backend and authentication method.
2. Connect securely to VPN server APIs.
3. Store users and subscriptions persistently.
4. Enforce traffic limits and expiration on the server.
5. Add access control and audit logs.

**Security:** Never put API keys, server passwords, or other secrets in frontend code.

---

## فارسی

### نصب و اجرای محلی
**روش اول: دانلود**
1. وارد همین مخزن در GitHub شوید.
2. روی **Code** و سپس **Download ZIP** بزنید.
3. فایل ZIP را روی گوشی یا رایانه استخراج کنید.
4. فایل `index.html` را با یک مرورگر به‌روز باز کنید.

**روش دوم: دریافت با Git**
```bash
git clone https://github.com/I2i1jwjkw/dark-panel.git
cd dark-panel
```
سپس فایل `index.html` را در مرورگر باز کنید. برای اجرای این نسخه نیازی به نصب پکیج یا فرایند build نیست.

### انتشار با GitHub Pages
1. وارد صفحه مخزن شوید و **Settings** را باز کنید.
2. از منوی کناری **Pages** را انتخاب کنید.
3. در بخش **Build and deployment** گزینه **Deploy from a branch** را انتخاب کنید.
4. شاخه `main` و پوشه `/(root)` را انتخاب کرده و **Save** را بزنید.
5. پس از انتشار، آدرس سایت در همان صفحه Pages نمایش داده می‌شود.

### امکانات
- داشبورد و آمار نمونه
- فهرست و جستجوی کاربران
- افزودن کاربر آزمایشی در همان صفحه
- بخش سرورها، ترافیک و تنظیمات
- طراحی مناسب موبایل و رایانه

### وضعیت فعلی
این پروژه فعلاً فقط **نمونه رابط کاربری (Frontend)** است. اطلاعات کاربران و آمار ساختگی هستند. افزودن کاربر فقط در حافظه همان صفحه انجام می‌شود و در دیتابیس ذخیره نمی‌شود. پنل هنوز به سرور VPN، Xray، بک‌اند یا API متصل نیست و برای استفاده عملیاتی آماده نیست.

### مراحل بعدی توسعه
1. انتخاب بک‌اند و روش احراز هویت
2. اتصال امن به API سرورهای VPN
3. ذخیره دائمی کاربران و اشتراک‌ها
4. اعمال محدودیت حجم و تاریخ انقضا در سمت سرور
5. افزودن کنترل دسترسی و ثبت رویدادها

**امنیت:** کلیدهای API، رمز سرور و اطلاعات محرمانه را داخل کد فرانت‌اند قرار ندهید.
