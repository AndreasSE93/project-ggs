package clientLobby;

import packageManaging.LobbyHandler;
import graphicalReference.LobbyGUI;

public class LobbyShell {
	
	public LobbyHandler handler;
	public LobbyGUI graphInterface;
	
	public LobbyShell() {
		this.handler = new LobbyHandler();
		this.graphInterface = new LobbyGUI();
	}
	
	public void Initiate() {
		
	}
	
	public void ShutDown() {
		
	}
	
}
