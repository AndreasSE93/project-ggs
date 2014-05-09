package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.CreateGameMessage;
import packageManaging.CreateGameMessageEncoder;
import packageManaging.HostRoom;
import packageManaging.JoinMessage;
import packageManaging.JoinMessageEncoder;
import packageManaging.LobbyMessageEncoder;
import packageManaging.LobbyServerMessage;
import packageManaging.Message;
import packageManaging.ChatMessageEncoder;
import packageManaging.RefeshMessageEncoder;
import packageManaging.RefreshMessage;
import graphicalReference.ChatGUI;
import graphicalReference.LobbyGUI;

import clientNetworking.NetManager;

public class LobbyHandler implements HandlerInterface,
		ActionListener {
	LobbyGUI lg;
	ChatGUI cg;
	LobbyMessageEncoder lme = new LobbyMessageEncoder();
	ChatMessageEncoder cme = new ChatMessageEncoder();
	JoinMessageEncoder jme = new JoinMessageEncoder();
	CreateGameMessageEncoder gme = new CreateGameMessageEncoder();
	RefeshMessageEncoder  rme = new RefeshMessageEncoder();
	NetManager network;
	boolean loop = true;
	int state;
	final String userName;
	public LobbyHandler(NetManager net, String username){
		this.network = net;
		this.userName = username;
	}
	
	public int init() {

		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = network.receiveMessage();
			LM = lme.decode(firstcall);
			
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}

		if (LM != null) {
			lg = new LobbyGUI(this.userName);
			lg.render(LM.UserList, LM.GameHost);
			lg.chatgui.field.addActionListener(this);
			lg.joinButton.addActionListener(this);
			lg.createButton.addActionListener(this);
			lg.refreshButton.addActionListener(this);
			
		} else {
			System.out.println("Connection failed");
		}

		while (loop) {
			try {
				String mess = recieveMessage();
			
				int id = retrieveId(mess);
				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
		
		//Change state!!!
		
		 lg.lobby.setVisible(false);
		 return state;

	}

	public void actionPerformed(ActionEvent e) {
		String command = e.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "chatmessage":  //JTextField
			try {
				JSONtext = cme.encode(new Message(lg.chatgui.field.getText(), lg.chatgui.userName));
				lg.chatgui.field.setText("");
			} catch (JSONException e1) {
				e1.printStackTrace();
			}

			break;

		case "joinbutton":
			//try {
			//lg.jt.setCellSelectionEnabled(true);
			lg.jt.setRowSelectionAllowed(true);
			System.out.println(lg.jt.getModel().getValueAt(1, 3));
				System.out.println(lg.jt.getSelectedRow() + " hej");
				/*JSONtext = jme.encode(new JoinMessage((HostRoom) ); 
				
				} catch (JSONException e1) {
					// TODO Auto-generated catch block
					e1.printStackTrace();
				}*/
			state = 2; //Temporärt för att byta till tic tac toe
			loop=false; // -||-_________________________________


			break;
			
		case "createbutton":
			try {
				
				JSONtext = gme.encode(new CreateGameMessage(lg.createList.getSelectedValue(), 2 , this.userName));
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			
			break;
		case "refreshbutton":			//SKicka jag vill uppdatera spelarlistorna, server skickar ut nya listor
			try {
				JSONtext = rme.encode (new RefreshMessage());
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			break;
		default:
			JSONtext = "";

		}
		sendMessage(JSONtext);

	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		case 100: // Message.java
			ChatMessageEncoder mh = new ChatMessageEncoder();
			Message chatMessage = mh.decode(message);
			lg.chatgui.chatUpdate(chatMessage.message, chatMessage.user);
			break;

		case 101: // Create session
			
			break;

		case 102: // Join session
			break;

		case 103: // LobbyClientMessage.java/LobbyServerMessage.java
			LobbyServerMessage ls = lme.decode(message);
			lg.addArrayList(lg.eastTopPanel, ls.UserList);
			break;
		default:
			System.out.println("something went wrong! Error code: " + id);
		}

	}

	public void sendMessage(String message) {
		network.send(message);
	}

	public String recieveMessage() throws IOException {
		return network.receiveMessage();

	}

	public int retrieveId(String mess) {
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;

	}
}
