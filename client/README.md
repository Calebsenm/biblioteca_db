# 🎨 Frontend - Sistema de Biblioteca

Frontend desarrollado en React para el sistema de gestión de biblioteca.

## 🏗️ Arquitectura

### Estructura del Proyecto
```
client/
├── public/                    # Archivos públicos
│   ├── index.html             # HTML principal
│   └── vite.svg               # Iconos
├── src/
│   ├── components/            # Componentes React
│   │   ├── Admin/             # Componentes de administrador
│   │   │   ├── Dashboard.jsx
│   │   │   ├── Books.jsx
│   │   │   ├── Users.jsx
│   │   │   ├── Loans.jsx
│   │   │   ├── Fines.jsx
│   │   │   ├── Reservation.jsx
│   │   │   ├── CreateBook.jsx
│   │   │   ├── CreateEditorial.jsx
│   │   │   └── Author.jsx
│   │   ├── User/                # Componentes de usuario
│   │   │   ├── Dashboard.jsx
│   │   │   ├── Libro.jsx
│   │   │   ├── Prestamos.jsx
│   │   │   ├── Reservas.jsx
│   │   │   ├── Multas.jsx
│   │   │   └── Historial.jsx
│   │   ├── Layout/              # Componentes de layout
│   │   │   ├── AdminLayout.jsx
│   │   │   ├── UserLayout.jsx
│   │   │   └── Sidebar.jsx
│   │   ├── common/              # Componentes comunes
│   │   │   ├── DataTable.jsx
│   │   │   ├── FilterForm.jsx
│   │   │   ├── Modal.jsx
│   │   │   ├── ConfirmDialog.jsx
│   │   │   ├── Notification.jsx
│   │   │   └── Breadcrumbs.jsx
│   │   └── Navbar.jsx
│   ├── pages/                   # Páginas principales
│   │   ├── Login.jsx
│   │   ├── Register.jsx
│   │   ├── NotFound.jsx
│   │   └── DashboardRedirect.jsx
│   ├── routes/                  # Configuración de rutas
│   │   ├── AdminRoutes.jsx
│   │   ├── UserRoutes.jsx
│   │   └── ProtectedRoute.jsx
│   ├── services/                # Servicios API
│   │   ├── api.js               # Cliente HTTP base
│   │   ├── authService.js
│   │   ├── bookService.js
│   │   ├── loanService.js
│   │   ├── fineService.js
│   │   ├── reservationService.js
│   │   ├── userService.js
│   │   ├── authorService.js
│   │   └── editorialService.js
│   ├── hooks/                   # Hooks personalizados
│   │   ├── useApi.js            # Hook para llamadas API
│   │   └── useNotification.js
│   ├── styles/                  # Estilos CSS
│   │   ├── common/              # Estilos comunes
│   │   ├── admin/               # Estilos de admin
│   │   └── user/                # Estilos de usuario
│   ├── assets/                  # Recursos estáticos
│   ├── config.js                # Configuración
│   ├── auth.js                  # Utilidades de autenticación
│   ├── App.jsx                  # Componente principal
│   └── main.jsx                 # Punto de entrada
├── package.json                 # Dependencias y scripts
├── vite.config.js               # Configuración de Vite
└── index.html                  # HTML principal
```

## 🚀 Tecnologías

### Core
- **React 18** - Framework de UI
- **Vite** - Build tool y dev server
- **React Router v6** - Navegación

### Estilos
- **CSS3** - Estilos nativos
- **CSS Modules** - Modularización de estilos

### Utilidades
- **Axios** - Cliente HTTP
- **React Hooks** - Estado y efectos

## 🔧 Configuración

### Variables de Entorno
El frontend usa variables de entorno de Vite para la configuración:

```env
# Archivo .env en client/
VITE_URL_BACKEND=http://localhost:4000/v1/api
```

### Configuración por Defecto
Si no se especifica `VITE_URL_BACKEND`, el frontend usa:
```
http://localhost:4000/v1/api
```

### Dependencias Principales
```json
{
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-router-dom": "^6.8.1"
  },
  "devDependencies": {
    "@vitejs/plugin-react": "^4.0.0",
    "vite": "^4.3.9"
  }
}
```

## 🚀 Ejecución

### Desarrollo
```bash
# Instalar dependencias
npm install

# Ejecutar servidor de desarrollo
npm run dev
```

### Producción
```bash
# Build para producción
npm run build

# Preview del build
npm run preview
```

### Scripts Disponibles
```bash
npm run dev        # Servidor de desarrollo
npm run build      # Build para producción
npm run preview    # Preview del build
npm run lint       # Linting del código
```

## 🎯 Funcionalidades

### Autenticación
- **Login/Logout** - Gestión de sesiones
- **Registro** - Creación de usuarios
- **JWT** - Tokens de autenticación
- **Protección de rutas** - Acceso basado en roles

### Dashboard de Usuario
- **Vista general** - Resumen de actividad
- **Libros disponibles** - Catálogo de libros
- **Préstamos activos** - Estado de préstamos
- **Reservas** - Gestión de reservas
- **Multas** - Consulta y pago de multas
- **Historial** - Historial de préstamos

### Dashboard de Administrador
- **Panel de control** - Estadísticas generales
- **Gestión de libros** - CRUD completo
- **Gestión de usuarios** - Administración de usuarios
- **Préstamos** - Vista de todos los préstamos
- **Multas** - Gestión de multas
- **Reservas** - Gestión de reservas
- **Editoriales y autores** - Gestión de catálogos

## 🔐 Sistema de Roles

### Usuario Normal
```javascript
// Permisos
- books:read          // Ver libros disponibles
- loans:create        // Crear préstamos
- loans:view          // Ver préstamos propios
- fines:read          // Ver multas propias
- fines:create        // Pagar multas
- reservations:create // Crear reservas
- reservations:view   // Ver reservas propias
```

### Administrador
```javascript
// Permisos adicionales
- books:write         // Crear/editar libros
- books:delete        // Eliminar libros
- loans:manage        // Gestionar préstamos
- users:view          // Ver usuarios
- users:manage        // Gestionar usuarios
- fines:read          // Ver todas las multas
```

## 🎨 Componentes

### Componentes de Layout
- **AdminLayout** - Layout para administradores
- **UserLayout** - Layout para usuarios
- **Sidebar** - Menú lateral con toggle
- **Navbar** - Barra de navegación

### Componentes Comunes
- **DataTable** - Tabla de datos con paginación
- **FilterForm** - Formulario de filtros
- **Modal** - Ventana modal reutilizable
- **ConfirmDialog** - Diálogo de confirmación
- **Notification** - Sistema de notificaciones
- **Breadcrumbs** - Navegación de migas

### Componentes de Admin
- **Dashboard** - Panel de control
- **Books** - Gestión de libros
- **Users** - Gestión de usuarios
- **Loans** - Gestión de préstamos
- **Fines** - Gestión de multas
- **Reservation** - Gestión de reservas

### Componentes de Usuario
- **Dashboard** - Vista general del usuario
- **Libro** - Catálogo de libros
- **Prestamos** - Préstamos del usuario
- **Reservas** - Reservas del usuario
- **Multas** - Multas del usuario
- **Historial** - Historial de préstamos

## 🔄 Hooks Personalizados

### useApi
```javascript
// Hook para llamadas API con estado
const { data, loading, error, refetch } = useApiGet(
  () => bookService.getBooks(),
  [dependencies]
);

const { execute, loading, error, success } = useApiMutation(
  (data) => bookService.createBook(data)
);
```

### useFilters
```javascript
// Hook para manejo de filtros
const { filters, updateFilter, clearFilters } = useFilters({
  titulo: "",
  genero: "",
  autor: ""
});
```

### useForm
```javascript
// Hook para manejo de formularios
const { formData, handleChange, resetForm } = useForm({
  titulo: "",
  genero: "",
  autor_id: ""
});
```

## 🌐 Servicios API

### Estructura de Servicios
```javascript
// Ejemplo: bookService.js
export const bookService = {
  // Obtener libros
  getBooks: async (params = {}) => {
    const response = await api.get('/books', { params });
    return response.data;
  },

  // Crear libro
  createBook: async (data) => {
    const response = await api.post('/admin/books', data);
    return response.data;
  },

  // Actualizar libro
  updateBook: async (id, data) => {
    const response = await api.post(`/admin/books/${id}`, data);
    return response.data;
  },

  // Eliminar libro
  deleteBook: async (id) => {
    const response = await api.delete(`/admin/books/${id}`);
    return response.data;
  }
};
```

### Cliente HTTP Base
```javascript
// api.js
import axios from 'axios';

export const api = axios.create({
  baseURL: import.meta.env.VITE_URL_BACKEND,
  timeout: 10000,
});

// Interceptor para agregar token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Interceptor para manejar errores
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);
```

## 🎨 Sistema de Estilos

### Organización
```
styles/
├── common/           # Estilos comunes
│   ├── Admin.css
│   ├── App.css
│   ├── Layout.css
│   ├── Modal.css
│   ├── Navbar.css
│   ├── Sidebar.css
│   └── ...
├── admin/            # Estilos específicos de admin
│   ├── Dashboard.css
│   ├── Books.css
│   ├── Users.css
│   └── ...
└── user/             # Estilos específicos de usuario
    ├── Dashboard.css
    ├── Libro.css
    ├── Prestamos.css
    └── ...
```

### Variables CSS
```css
:root {
  --primary-color: #10b981;
  --secondary-color: #3b82f6;
  --danger-color: #ef4444;
  --warning-color: #f59e0b;
  --success-color: #10b981;
  --text-color: #333;
  --bg-color: #f9fafb;
  --border-color: #e5e7eb;
}
```

## 🔄 Estado y Gestión de Datos

### Estado Local
- **useState** - Estado de componentes
- **useEffect** - Efectos secundarios
- **useContext** - Estado global (si es necesario)

### Estado de API
- **useApiGet** - Datos de lectura
- **useApiMutation** - Operaciones de escritura
- **Loading states** - Estados de carga
- **Error handling** - Manejo de errores

## 🛡️ Seguridad

### Autenticación
- **JWT Tokens** - Almacenados en localStorage
- **Protección de rutas** - Basada en roles
- **Interceptores** - Manejo automático de tokens

### Validación
- **Validación de formularios** - Cliente y servidor
- **Sanitización** - Limpieza de datos
- **CSRF Protection** - Headers de seguridad

## 📱 Responsive Design

### Breakpoints
```css
/* Mobile First */
@media (min-width: 768px) { /* Tablet */ }
@media (min-width: 1024px) { /* Desktop */ }
@media (min-width: 1280px) { /* Large Desktop */ }
```

### Componentes Responsive
- **Sidebar** - Colapsable en móvil
- **DataTable** - Scroll horizontal en móvil
- **Modal** - Full screen en móvil
- **Forms** - Stack vertical en móvil

## 🧪 Testing

### Configuración de Tests
```bash
# Instalar dependencias de testing
npm install --save-dev @testing-library/react @testing-library/jest-dom

# Ejecutar tests
npm test

# Tests con coverage
npm run test:coverage
```

### Ejemplos de Tests
```javascript
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Dashboard from './Dashboard';

test('renders dashboard title', () => {
  render(
    <BrowserRouter>
      <Dashboard />
    </BrowserRouter>
  );
  expect(screen.getByText(/Panel de Administración/i)).toBeInTheDocument();
});
```

## 🚀 Optimización

### Performance
- **Code Splitting** - Carga lazy de componentes
- **Memoization** - React.memo para componentes
- **Bundle Analysis** - Análisis de tamaño de bundle

### Build Optimization
- **Tree Shaking** - Eliminación de código no usado
- **Minification** - Compresión de código
- **Asset Optimization** - Optimización de imágenes

## 🔧 Desarrollo

### Convenciones
- **PascalCase** - Componentes React
- **camelCase** - Variables y funciones
- **kebab-case** - Archivos CSS
- **snake_case** - Variables de entorno

### Estructura de Commits
```
feat: add user dashboard
fix: resolve login issue
docs: update README
style: improve button styling
refactor: simplify component logic
test: add unit tests for auth
```

## 🚀 Despliegue

### Build de Producción
```bash
npm run build
```

### Variables de Producción
```env
VITE_URL_BACKEND=https://api.tubiblioteca.com/v1/api
```


