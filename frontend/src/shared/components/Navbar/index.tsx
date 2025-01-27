import React from 'react';
import { WeManageIt } from 'shared/icons';
import {
  Container,
  LogoWrapper,
  IconWrapper,
  Logo,
  LogoTitle,
  ActionContainer,
  ActionButtonContainer,
  ActionButtonWrapper,
  ActionButtonTitle,
} from './Styles';

type ActionButtonProps = {
  name: string;
  active?: boolean;
};

export const ActionButton: React.FC<ActionButtonProps> = ({ name, active, children }) => {
  return (
    <ActionButtonWrapper active={active ?? false}>
      <IconWrapper>{children}</IconWrapper>
      <ActionButtonTitle>{name}</ActionButtonTitle>
    </ActionButtonWrapper>
  );
};

export const ButtonContainer: React.FC = ({ children }) => (
  <ActionContainer>
    <ActionButtonContainer>{children}</ActionButtonContainer>
  </ActionContainer>
);

export const PrimaryLogo = () => {
  return (
    <LogoWrapper>
      <WeManageIt width={42} height={42} />
      <LogoTitle>WeManageIt</LogoTitle>
    </LogoWrapper>
  );
};

const Navbar: React.FC = ({ children }) => {
  return <Container>{children}</Container>;
};

export default Navbar;
