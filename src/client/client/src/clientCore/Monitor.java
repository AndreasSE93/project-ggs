package clientCore;

import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Calendar;

import org.json.JSONException;

import packageManaging.LobbyServerMessage;
import packageManaging.Message;
import packageManaging.MessageHandler;

import clientLobby.LobbyShell;
import clientNetworking.NetManager;
import clientNetworking.Connection;
import clientChat.ChatShell;

public class Monitor {
	
	Connection conn;
	NetManager net;
	ChatShell chatModule;
	LobbyShell lobbyModule;
	
	public Monitor() {
		this.conn = new Connection("130.238.95.70", 8080);
		this.chatModule = new ChatShell();
		this.lobbyModule = new LobbyShell();
		this.net = new NetManager(this.chatModule.handler, this.conn);
		
		try {
			this.net.connectToServer();
		} catch (IOException e) {
			System.out.println("Connection to server failed!\n");
			e.printStackTrace();
		}
		this.runner();
	}
	
	private void runner() {
		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = net.receiveMessage();
			System.out.println(firstcall + "\n");
			LM = this.lobbyModule.handler.decode(firstcall);
			System.out.println(LM + "\n");
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}
		if (LM != null) {
			this.lobbyModule.Initiate(LM.UserList, LM.GameHost, this.net);
		} else {
			System.out.println("Connection failed");
		}
		
		String mess;
		LobbyServerMessage LSM = null;
		
		while(true) {
			try {
				MessageHandler MH = new MessageHandler();
				mess = net.receiveMessage();
				Message m = MH.decode(mess);
				this.lobbyModule.graphInterface.chatPanel.setText(this.lobbyModule.graphInterface.chatPanel.getText() + "<" + getTime() + "> " + m.user + ": "  + m.message + "\n");
				this.lobbyModule.graphInterface.chatPanel.setCaretPosition(this.lobbyModule.graphInterface.chatPanel.getDocument().getLength());
			} catch (IOException | JSONException e) {
			
				e.printStackTrace();
			}
		}
	}
	
	public String getTime() {
    	Calendar cal = Calendar.getInstance();
    	cal.getTime();
    	SimpleDateFormat sdf = new SimpleDateFormat("HH:mm:ss");
    	return sdf.format(cal.getTime());
    }

	
}
