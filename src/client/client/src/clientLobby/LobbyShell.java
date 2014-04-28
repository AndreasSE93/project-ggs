package clientLobby;

import java.util.ArrayList;

import clientNetworking.NetManager;

import packageManaging.LobbyHandler;
import graphicalReference.LobbyGUI;

public class LobbyShell {
	
	public LobbyHandler handler;
	public LobbyGUI graphInterface;
	
	public LobbyShell() {
		this.handler = new LobbyHandler();
		this.graphInterface = new LobbyGUI();
	}
	
	public void Initiate(ArrayList<String> players,  ArrayList<String> games, NetManager nm) {
		this.graphInterface.render(players, games, nm);
	}
	
	public void ShutDown() {
		
	}
	
}
