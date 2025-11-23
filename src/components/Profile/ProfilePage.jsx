import React from 'react';
import { auth } from '../../lib/firebase';

export default function ProfilePage() {
  const user = auth.currentUser;
  return (
    <div className="container">
      <h2>My Vibe</h2>
      <p>Username: {user?.email}</p>
      <p>GruvCoin: 1,250 ₲ (demo)</p>
      <button onClick={() => auth.signOut()}>Sign Out</button>
    </div>
  );
}