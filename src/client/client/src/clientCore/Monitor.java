package clientCore;

import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Calendar;

import javax.swing.JOptionPane;

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
<<<<<<< HEAD
		this.conn = new Connection("localhost", 8080);
=======
		String serverConnectStr = "Enter server: ";
		while (true) {
			String hostStr = JOptionPane.showInputDialog(serverConnectStr, Connection.DEFAULT_HOST + ":" + String.valueOf(Connection.DEFAULT_PORT));
			if (hostStr == null) {
				return;
			}
			try {
				this.conn = new Connection(hostStr);
				break;
			} catch (NumberFormatException e) {
				serverConnectStr = "Invalid port. Enter server:";
			}
		}
		
>>>>>>> eaeef93d90fa8e0a26e42171122e861731b88f87
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
			this.chatModule.Initiate(this.lobbyModule.graphInterface.chatt, this.net, this.lobbyModule.graphInterface.lobby);
		} else {
			System.out.println("Connection failed");
		}
		
		String mess;
		
		
		while(true) {
			try {
				
				MessageHandler MH = new MessageHandler();
				mess = net.receiveMessage();
				Message m = MH.decode(mess);
				this.chatModule.graphInterface.chatPanel.setText(this.chatModule.graphInterface.chatPanel.getText() + "<" + getTime() + "> " + m.user + ": "  + m.message + "\n");
				this.chatModule.graphInterface.chatPanel.setCaretPosition(this.chatModule.graphInterface.chatPanel.getDocument().getLength());
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
