import React, { useState } from 'react';
import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../../lib/firebase';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [err, setErr] = useState('');

  const handleSignIn = async (e) => {
    e.preventDefault();
    setErr('');
    try { await signInWithEmailAndPassword(auth, email, password); } catch (e) { setErr(e.message); }
  };

  const handleSignUp = async (e) => {
    e.preventDefault();
    setErr('');
    try { await createUserWithEmailAndPassword(auth, email, password); } catch (e) { setErr(e.message); }
  };

  return (
    <div className="centered">
      <div className="card">
        <h1>The Gruv</h1>
        <input placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
        <input placeholder="Password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        {err && <div className="error">{err}</div>}
        <div className="row">
          <button onClick={handleSignIn}>Log In</button>
          <button onClick={handleSignUp}>Sign Up</button>
        </div>
      </div>
    </div>
  );
}