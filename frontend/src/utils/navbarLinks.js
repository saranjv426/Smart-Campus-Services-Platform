export const getRoleBasedLinks = (isLoggedIn, currentUser) => {
  const baseLinks = [
    { label: 'Home', path: '/' },
    { label: 'Services', path: '/services' },
  ];

  if (!isLoggedIn || !currentUser) {
    return {
      navLinks: baseLinks,
      authLinks: [
        { label: 'Login', path: '/login' },
        { label: 'Register', path: '/register', className: 'register-btn' },
      ],
    };
  }

  const roleLinks = {
    student: [{ label: 'My Bookings', path: '/bookings' }],
    staff: [{ label: 'Staff Dashboard', path: '/dashboard/staff' }],
    admin: [{ label: 'Admin Dashboard', path: '/dashboard/admin' }],
  };

  const authLinks = [
    ...(currentUser.role === 'student'
      ? [{ label: 'Profile', path: `/profile/${currentUser.id}` }]
      : []),
  ];

  return {
    navLinks: [...baseLinks, ...(roleLinks[currentUser.role] || [])],
    authLinks,
  };
};