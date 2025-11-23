import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from './components/Auth/LoginPage';
import Feed from './components/Feed/Feed';
import EventDetail from './components/Event/EventDetail';
import ProfilePage from './components/Profile/ProfilePage';
import CrewPage from './components/Crew/CrewPage';
import VibeVaultPage from './components/Vaults/VibeVaultPage';
import BurnerWallet from './components/Wallet/BurnerWallet';
import GruvDropClaim from './components/Drop/GruvDropClaim';
import { useAuthGuard } from './lib/useAuth';
import GlobalStyles from './styles/GlobalStyles';

export default function App() {
  const { user, loading } = useAuthGuard();

  if (loading) return <div className="centered">Authenticating...</div>;
  if (!user) return <LoginPage />;

  return (
    <>  
      <GlobalStyles />
      <Routes>
        <Route path="/" element={<Feed />} />
        <Route path="/event/:id" element={<EventDetail />} />
        <Route path="/profile" element={<ProfilePage />} />
        <Route path="/crew" element={<CrewPage />} />
        <Route path="/vaults" element={<VibeVaultPage />} />
        <Route path="/wallet" element={<BurnerWallet />} />
        <Route path="/drop/:id" element={<GruvDropClaim />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </>
  );
}